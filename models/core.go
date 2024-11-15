package models

import "pkg/dao"

// SQL 语句
type Core struct{}

func (Core) GetAllBlog() ([]Blog, error) {
	var blogs []Blog
	err := dao.MysqlClient.Raw("SELECT id, user_id, title, context, `like`, tags, created_by, collection, created_at, updated_at, deleted_at FROM blog ORDER BY created_at DESC;").Scan(&blogs).Error
	return blogs, err
}
