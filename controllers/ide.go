package controllers

import (
	"github.com/gin-gonic/gin"
	"pkg/utils/convert"
	"pkg/utils/request"
	"fmt"
	
)

type IDEController struct{}
type IDEJudgestruct struct {
	Code        string `json:"code"`
	Language    string `json:"language"`
	Input       string `json:"input"`
	TimeLimit   uint   `json:"timeLimit"`
	MemoryLimit uint   `json:"memoryLimit"`
	Args        string `json:"args"`
}

func (IDEController) JudgeCode(c *gin.Context) {
	ideJudgestruct := &IDEJudgestruct{}
	_ = c.BindJSON(&ideJudgestruct)
	str, _ := convert.Encode2String(ideJudgestruct)
	ip := AddrQueue.GetOne()
	resq, err := req.POST(fmt.Sprintf("%s/judge",ip), str)
	
	if err != nil {
		ReturnError(c, 1, err)
		return
	}
	returnstruct := &RetuenStruct{}
	_ = convert.Decode2Object(resq, &returnstruct)
	if returnstruct.Err != "" {
		ReturnError(c, 1, returnstruct.Err)
		return
	}
	ReturnSuccess(c, 0, "success", returnstruct.Info, 1)
}
