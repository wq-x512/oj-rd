package models

import (
	"pkg/dao"
	"time"
)

type ChatRecord struct {
	SenderID   string
	ReceiverID string
	Status     uint16 // 0 未读 1 已读
	Content    string
	CreatedAt  time.Time
	DeletedAt  time.Time
}

func CreateChatRecord(chatRecord *ChatRecord) error {
	return dao.InsertDocument("chat", "chat_record", chatRecord)
}
