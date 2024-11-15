package models

import (
	"pkg/dao"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID         string  `json:"id" gorm:"primarykey"`        // 题目id
	UserID     string  `json:"userID"`                      // 用户id
	Title      string  `json:"title"`                       // 标题
	Context    string  `json:"context"`                     //文本
	Tags       *string `json:"tags"`                        //标签
	Collection int32   `json:"collection" gorm:"default:0"` // 收藏
	Like       int32   `json:"like" gorm:"default:0"`       //喜欢
	CreatedBy  string  `json:"createdBy"`                   // 用户名
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (Blog) TableName() string {
	return "blog"
}
func CreateBlog(blog *Blog) error {
	err := dao.MysqlClient.Create(blog).Error
	return err
}
func ChangeBlog(blog *Blog) error {
	err := dao.MysqlClient.Updates(blog).Error
	return err
}
func QueryBlog(id string) (Blog, error) {
	var blog Blog
	err := dao.MysqlClient.Where("id = ?", id).First(&blog).Error
	return blog, err
}

func (Blog) GetAllBlog() ([]Blog, error) {
	var blogs []Blog
	err := dao.MysqlClient.Order("created_at DESC").Find(&blogs).Error
	return blogs, err
}
