package models

import "pkg/dao"

type Contest struct {
	ID string `json:"id" gorm:"primarykey"`
}

func (Contest) TableName() string {
	return "contest"
}
func CreateContest(contest *Contest) error {
	err := dao.MysqlClient.Create(contest).Error
	return err
}
func QueryContest(id string) (Contest, error) {
	var contest Contest
	err := dao.MysqlClient.Where("id = ?", id).First(&contest).Error
	return contest, err
}

func GetAllContest() ([]Contest, error) {
	var contests []Contest
	err := dao.MysqlClient.Find(&contests).Error
	return contests, err
}
