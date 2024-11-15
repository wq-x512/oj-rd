package models

import (
	"pkg/dao"
	"time"

	"gorm.io/gorm"
)

type Training struct {
	ID           string  `json:"id" gorm:"primaryKey"`
	Introduction *string `json:"introduction"`
	CreatedBy    string  `json:"createdBy"`
	UserID       string  `json:"userID"`
	IsPrivate    bool    `json:"isPrivate"`
	Like         int32   `json:"like"`
	Collection   int32   `json:"collection"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
type TrainingToProblem struct {
	TrainingID string `json:"trainingID"`
	ProblemID  string `json:"problemID"`
	Submission int32  `json:"submission"`
	Accept     int32  `json:"accept"`
}

func (Training) TableName() string {
	return "training"
}
func (TrainingToProblem) TableName() string {
	return "training_to_problem"
}

func  CreateTraining(training *Training) error {
	err := dao.MysqlClient.Create(training).Error
	return err
}
func  QueryTraining(id string) (Training, error) {
	var training Training
	err := dao.MysqlClient.Where("id = ?", id).First(&training).Error
	return training, err
}

func  GetAllTraining() ([]Training, error) {
	var trainings []Training
	err := dao.MysqlClient.Find(&trainings).Error
	return trainings, err
}
