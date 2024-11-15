package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pkg/models"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatController struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageStruct struct {
	Timestamp   int64   `json:"timestamp" gorm:"primaryKey;autoIncrement"`
	CurrentTime int64   `json:"currenttime" gorm:"autoUpdateTime:nano"`
	Sender      string  `json:"sender"`
	Receiver    string  `json:"receiver"`
	Text        *string `json:"text"`
	Type        uint    `json:"type"` // 1私信 2群发 3心跳检测
}

type Client struct {
	Conn          *websocket.Conn
	Addr          string
	HeartBeatTime uint64
	DataQueue     chan []byte
	GroupSets     sync.Map
}

var clients = make(map[string]*Client)
var clientsMu = &sync.Mutex{}

// var logger *log.Logger
// var file *os.File

func (ChatController) Handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	sender_id := c.Query("id")
	currentTime := uint64(time.Now().Unix())
	client := &Client{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(),
		HeartBeatTime: currentTime,
		DataQueue:     make(chan []byte, 1024),
	}
	clientsMu.Lock()
	clients[sender_id] = client
	clientsMu.Unlock()
	go client.sendProc()
	go client.recvProc()
	// go client.checkConnect(sender_id)
	client.Conn.WriteMessage(websocket.TextMessage, []byte("Hello Client!"))
}

func findClient(id string) *Client {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	return clients[id]
}
func (client *Client) checkConnect(sender_id string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("函数内部发生 panic: %v", r)
		}
	}()
	for {
		_, _, err := client.Conn.NextReader()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) ||
				errors.Is(err, io.EOF) || websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("WebSocket 连接已关闭")
			} else {
				fmt.Printf("NextReader 遇到未知错误: %v", err)
			}
			clientsMu.Lock()
			delete(clients, sender_id)
			clientsMu.Unlock()
			if err := client.Conn.Close(); err != nil {
				fmt.Printf("关闭 WebSocket 连接失败: %v", err)
			}
			break
		}
	}
}
func (client *Client) sendProc() {
	for {
		select {
		case data := <-client.DataQueue:
			fmt.Println("[ws] sendProc >>> msg :", string(data))
			err := client.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("发送消息错误")
				continue
			}

		}
	}
}
func (client *Client) recvProc() {
	for {
		fmt.Println("等待消息")
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("传输错误", err)
			return
		}
		fmt.Println("消息 :", message)
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			fmt.Printf("error: %v \n", err)
			break
		}

		var msg MessageStruct
		err = json.Unmarshal([]byte(message), &msg)
		if err != nil {
			fmt.Println("解码失败 :", err)
			continue
		}
		/////////////////////////////////
		go func() {
			chat_record := &models.ChatRecord{
				SenderID:   msg.Sender,
				ReceiverID: msg.Receiver,
				Status:     0,
				Content:    *msg.Text,
				CreatedAt:  time.Now(),
			}
			err := models.CreateChatRecord(chat_record)
			if err != nil {
				fmt.Println("插入出错", err)
			} else {
				fmt.Println("插入成功")
			}
		}()
		////////////////////////////////
		dispatch(msg)
		fmt.Println("[ws] recvProc <<< ", msg)
	}
	fmt.Println("消息函数关闭")
}
func dispatch(msg MessageStruct) {
	fmt.Println("消息分发中")
	switch msg.Type {
	case 1:
		sendMsg(msg)
	case 2:
		sendGroupMsg(msg)
	case 3:
		HeartBeatMsg(msg)
	default:
		fmt.Printf("未知消息类型: %d", msg.Type)
	}
}
func broadMsg(msg MessageStruct) error { // 局域网广播
	return nil
}
func HeartBeatMsg(msg MessageStruct) { // 心跳检测
	fmt.Printf("心跳检测[%s]来自%s \n", *msg.Text, msg.Sender)
	target_client := findClient(msg.Receiver)
	if target_client != nil {
		fmt.Printf("准备向目标客户端发送消息: %s \n", *msg.Text)
		respond := "pong"
		msg.Text = &respond
		message, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("编码失败 : ", err)
		}
		target_client.DataQueue <- []byte(message)
	}
}
func sendMsg(msg MessageStruct) { //私聊
	fmt.Printf("私发消息[%s]给%s \n", *msg.Text, msg.Receiver)
	target_client := findClient(msg.Receiver)
	if target_client != nil {
		fmt.Printf("准备向目标客户端发送消息: %s \n", *msg.Text)
		message, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("编码失败 : ", err)
		}
		target_client.DataQueue <- []byte(message)
	} else {
		fmt.Printf("目标客户端 %s 未找到或者未建立连接 \n", msg.Receiver)
	}
}
func sendGroupMsg(msg MessageStruct) { //群聊
	data := *msg.Text
	target_id := msg.Receiver
	fmt.Printf("群发消息[%s]到%s", data, target_id)
}

// func initLogger() (*log.Logger, *os.File, error) {
// 	logFileName := "../server.log"
// 	if _, err := os.Stat(logFileName); os.IsExist(err) {
// 		if err := os.Remove(logFileName); err != nil {
// 			return nil, nil, fmt.Errorf("删除现有日志文件失败: %w", err)
// 		}
// 	}
// 	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("打开或创建日志文件失败: %w", err)
// 	}
// 	logger := log.New(file, "", log.LstdFlags)
// 	return logger, file, nil
// }
// func (ChatController) init() {
// 	var err error
// 	logger, file, err = initLogger()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "初始化日志失败：%v", err)
// 		os.Exit(1)
// 	}
// 	defer file.Sync()
// }
