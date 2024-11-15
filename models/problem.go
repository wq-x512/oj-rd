package models

import (
	"pkg/dao"
	"time"

	"gorm.io/gorm"
)

type Problem struct {
	ID                string  `json:"id" gorm:"primarykey"`
	Title             string  `json:"title"`
	Context           string  `json:"context"`
	InputDescription  string  `json:"inputDescription"`
	OutputDescription string  `json:"outputDescription"`
	Tip               string  `json:"tip"`
	Difficulty        float32 `json:"difficulty"`
	JudgeCase         string  `json:"judgeCase"`
	JudgeConfig       string  `json:"judgeConfig"`
	Tags              *string `json:"tags"`
	Submission        int32   `json:"submission"`
	Accept            int32   `json:"accept"`
	Collection        int32   `json:"collection"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	CreatedBy         string         `json:"createdBy"`
}

func (Problem) TableName() string {
	return "problem"
}

func  CreateProblem(problem *Problem) error {
	err := dao.MysqlClient.Create(problem).Error
	return err
}
func  QueryProblem(id string) (Problem, error) {
	var problem Problem
	err := dao.MysqlClient.Where("id = ?", id).First(&problem).Error
	return problem, err
}

func  GetAllProblem() ([]Problem, error) {
	var problems []Problem
	err := dao.MysqlClient.Order("created_at DESC").Find(&problems).Error
	return problems, err
}
