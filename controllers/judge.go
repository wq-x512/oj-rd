package controllers

import (
	"fmt"
	"pkg/utils/convert"
	"pkg/utils/request"
)

type JudgeController struct{}

type JudgeStruct struct {
	ID          string `json:"id"` // problem id
	Account     string `json:"account"`
	UserID      string `json:"userID"`
	Code        string `json:"code"`
	Language    string `json:"language"`
	Input       string `json:"input"`
	TimeLimit   uint   `json:"timeLimit"`
	MemoryLimit uint   `json:"memoryLimit"`
	Args        string `json:"args"`
}

type RetuenStruct struct {
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
	Info interface{} `json:"info"`
}

func (JudgeController) EvaluateCode(judgeStruct JudgeStruct, ip string, channel chan RetuenStruct) {
	str, _ := convert.Encode2String(judgeStruct)
	resq, err := req.POST(fmt.Sprintf("%s%s", ip, "/judge"), str)
	if err != nil {
		return
	}
	returnstruct := RetuenStruct{}
	err = convert.Decode2Object(resq, &returnstruct)
	if err != nil {
		return
	}
	channel <- returnstruct
}

type RemoteJudge struct{}

func SendCode() {

}
