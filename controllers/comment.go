package controllers

import (
	"pkg/models"
	"pkg/utils/convert"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController struct{}

func (CommentController) GetComment(c *gin.Context) {
	postID := c.Query("id")
	var comments []models.Comment
	comments, err := models.QueryComments(postID)
	results := []interface{}{}
	for _, v := range comments {
		result, _ := convert.Object2Dict(v)
		result["subcomment"], _ = models.QuerySubComments(v.ID)
		results = append(results, result)
	}
	if err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", results, int64(len(results)))
}

func (CommentController) CreateComment(c *gin.Context) {
	commmet := &models.Comment{}
	_ = c.Bind(&commmet)
	commmet.ID = uuid.New().String()
	err := models.CreateComment(commmet)
	if err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", commmet, 1)

}
func (CommentController) CreateSubComment(c *gin.Context) {
	subcommmet := &models.SubComment{}
	_ = c.Bind(&subcommmet)
	subcommmet.ID = uuid.New().String()
	err := models.CreateSubComment(subcommmet)
	if err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", subcommmet, 1)
}
