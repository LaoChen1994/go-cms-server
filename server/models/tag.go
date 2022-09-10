package models

import (
	"fmt"
	"time"
)

type Tag struct {
	Model
	ID        uint   `gorm:"primary key" gorm:"column id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	ParentId  int    `gorm:"column:parent_id" json:"parent_id"`
	CreatedId uint   `gorm:"column:created_id" json:"created_id" `
	State     uint   `gorm:"column:state" json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	if err := DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error; err != nil {
		fmt.Println("查询标签错误", err)
	}
	return
}

func GetTagTotal(maps interface{}) (count *int64) {
	if err := DB.Model(&Tag{}).Where(maps).Count(count).Error; err != nil {
		fmt.Println("标签数量查询错误", err)
	}

	return
}

func ExistTagByName(tag *Tag) (exist bool) {
	exist = false

	DB.Select("id").Where("name = ?", tag.Name).First(&tag)

	if tag.ID > 0 {
		exist = true
		return
	}

	return
}

func ExistTagById(id int64) bool {
	var tag Tag

	if id <= 0 {
		return false
	}

	res := DB.Where("id = ?", id).First(&tag)
	if res.Error == nil {
		if res.RowsAffected > 0 {
			return true
		}
	}

	return false
}

func CreateTag(tag *Tag) (err error) {
	err = nil

	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()

	res := DB.Omit("ID").Create(&tag)

	if err = res.Error; err != nil {
		return
	}

	return
}

func UpdateTag(tag *Tag) (effect int64, err error) {
	tag.UpdatedAt = time.Now()
	result := DB.Model(&tag).Where("id = ?", tag.ID).Updates(&tag)

	err = result.Error

	if err != nil {
		return 0, err
	}

	effect = result.RowsAffected

	return
}

func DeleteTag(id int64) (effectRow int64, err error) {
	result := DB.Delete(&Tag{}, id)

	err = result.Error
	effectRow = result.RowsAffected

	if err != nil {
		return
	}

	return effectRow, nil
}
