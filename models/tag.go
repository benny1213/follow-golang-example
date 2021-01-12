package models

import "fmt"

// Tag : 数据库中tag的结构
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// GetTags : 分页获取tags
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	fmt.Printf("db: %v", db)
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTotal : 获取总条目
func GetTotal(maps interface{}) int {
	var count int
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}
