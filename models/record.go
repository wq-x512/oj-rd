package models

import (
	"pkg/dao"
	"time"

	"gorm.io/gorm"
)

type Result struct {
	Error   int     `json:"error"`
	Input   string  `json:"input"`
	Output  string  `json:"output"`
	Time    string  `json:"time"`
	Memory  float64 `json:"memory"`
	Display string  `json:"display"`
}
type Record struct {
	ID           string `json:"id" gorm:"primarykey"`
	UserId       string `json:"userId" gorm:"index"`
	UserName     string `json:"userName"`
	ProblemId    string `json:"problemId" gorm:"index"`
	ProblemTitle string `json:"problemTitle"`
	Code         string `json:"code"`
	Language     string `json:"language"`
	JudgeResult  string `json:"judgeResult"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Record) TableName() string {
	return "record"
}
func  CreateRecord(record *Record) error {
	err := dao.MysqlClient.Create(record).Error
	return err
}

func  QueryRecord(record Record) (Record, error) {
	err := dao.MysqlClient.Find(&record, record).Error
	return record, err
}
func  GetAllRecord() ([]Record, error) {
	var records []Record
	err := dao.MysqlClient.Order("created_at DESC").Find(&records).Error
	return records, err
}
