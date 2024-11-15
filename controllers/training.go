package controllers

import (
	"github.com/gin-gonic/gin"
	"pkg/models"
)

type TrainingController struct{}

func (TrainingController) GetList(c *gin.Context) {
	results, err := models.GetAllTraining()
	if err == nil {
		ReturnSuccess(c, 0, "success", results, int64(len(results)))
		return
	}
	ReturnError(c, 1, err)
}
