package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Tag : 数据库中tag的结构
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate ： 创建时添加时间戳标记
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate ： 更新时添加时间戳标记
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// GetTags : 分页获取tags
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTotal : 获取总条目
func GetTotal(maps interface{}) int {
	var count int
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}

// ExistTagByName : 检查tagname是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// ExistTagByID : 根据id检查tag是否存在
func ExistTagByID(id int) bool {
	var tag Tag
	db.Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag : 添加tag
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// DeleteTag : 删除tag
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// EditTag : 编辑tag
func EditTag(data Tag) bool {
	db.Model(&Tag{}).Where("id = ?", data.ID).Update(data)
	return true
}
