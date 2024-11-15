package models

import (
	"pkg/dao"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string     `json:"id" gorm:"primarykey"`
	Account       string     `json:"account" gorm:"unique"`
	Password      string     `json:"password"`
	Email         *string    `json:"email"`
	FirstName     *string    `json:"firstName"`
	LastName      *string    `json:"lastName"`
	Introduction  *string    `json:"introduction"`
	Rating        int16      `json:"rating" gorm:"default:1000"`
	School        string     `json:"school"`
	Avatar        string     `json:"avatar"`
	Education     string     `json:"education"`
	UserRole      string     `json:"userRole"`
	Gender        bool       `json:"gender"`
	Submisson     int32      `json:"submisson"`
	Accept        int32      `json:"accept"`
	Codeforces    string     `json:"codeforces"`
	Birthday      *time.Time `json:"birthday"`
	BannedEndTime *time.Time `json:"bannedEndTime"`
	IsDelete      bool       `json:"isDelete"`
	IsBanned      bool       `json:"isBanned"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(user User) error {
	err := dao.MysqlClient.Create(user).Error
	return err
}
func (user User) QueryUser() (User, error) {
	err := dao.MysqlClient.Find(&user, user).Error
	return user, err
}
func GetAllUsers() ([]User, error) {
	var users []User
	err := dao.MysqlClient.Find(&users).Error
	return users, err
}
