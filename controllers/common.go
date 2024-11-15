package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonSuccStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Info  interface{} `json:"info"`
	Count int64       `json:"count"`
}
type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, info interface{}, count int64) {
	json := &JsonSuccStruct{Code: code, Msg: msg, Info: info, Count: count}
	c.JSON(http.StatusOK, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}
