package controllers

import (
	"github.com/gin-gonic/gin"
	"pkg/models"
)

type RecodeController struct{}

func (RecodeController) GetRecodeInfo(c *gin.Context) {
	id := c.Param("id")
	record := models.Record{ID: id}
	record, err := models.QueryRecord(record)
	if err == nil {
		ReturnSuccess(c, 0, "success", record, 1)
		return
	}
	ReturnError(c, 1, err)
}
func (RecodeController) GetList(c *gin.Context) {
	results, err := models.GetAllRecord()
	if err == nil {
		ReturnSuccess(c, 0, "success", results, int64(len(results)))
		return
	}
	ReturnError(c, 1, err)
}
func (RecodeController) CreateRecord(c *gin.Context) {
	data := &models.Record{}
	_ = c.BindJSON(&data)
}
