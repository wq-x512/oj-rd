package controllers

import (
	"pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogController struct{}

func (BlogController) CreateBlog(c *gin.Context) {
	data := &models.Blog{}
	if Err := c.BindJSON(&data); Err != nil {
		ReturnError(c, 1, Err)
		return
	}
	data.ID = uuid.New().String()
	if err := (models.CreateBlog(data)); err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", data, 1)
}
func (BlogController) ChangeBlog(c *gin.Context) {
	data := &models.Blog{}
	if Err := c.BindJSON(&data); Err != nil {
		ReturnError(c, 1, Err)
		return
	}
	if err := (models.ChangeBlog(data)); err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", data, 1)
}
func (BlogController) DeleteBlog(c *gin.Context) {
	id := c.Query("id")
}
func (BlogController) GetBlogInfo(c *gin.Context) {
	id := c.Param("id")
	blog, err := models.QueryBlog(id)
	if err == nil {
		ReturnSuccess(c, 0, "suceess", blog, 1)
		return
	}
	ReturnError(c, 1, err)
}
func (BlogController) GetList(c *gin.Context) {
	results, err := models.Core{}.GetAllBlog()
	if err == nil {
		ReturnSuccess(c, 0, "success", results, int64(len(results)))
		return
	}
	ReturnError(c, 1, err)
}
