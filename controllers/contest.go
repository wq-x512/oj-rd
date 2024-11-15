package controllers

import (
	"github.com/gin-gonic/gin"
	"pkg/models"
)

type ContestController struct{}

func (ContestController) GetList(c *gin.Context) {
	results, err := models.GetAllContest()
	if err == nil {
		ReturnSuccess(c, 0, "success", results, int64(len(results)))
		return
	}
	ReturnError(c, 1, err)
}
