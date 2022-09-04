package models

import (
	"fmt"
	"pd-go-server/pkg/setting"
	"time"
)

type Tag struct {
	Model
	ID        uint   `gorm:"primary key" gorm:"column id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	ParentId  int    `gorm:"column:parent_id" json:"parent_id"`
	CreatedId uint   `gorm:"column:author" json:"author" `
	State     uint   `gorm:"column:state" json:"state"`
}

func (Tag) TableName() string {
	return fmt.Sprintf("%s%s", setting.DatabaseConf.Prefix, "tag")
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	if err := DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error; err != nil {
		fmt.Println("查询标签错误", err)
	}
	return
}

func GetTagTotal(maps interface{}) (count int) {
	if err := DB.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		fmt.Println("标签数量查询错误", err)
	}

	return
}

func ExistTagByName(name string) bool {
	var tag Tag

	DB.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistTagById(id int64) bool {
	var tag Tag
	res := DB.Where("id = ?", id).First(&tag)
	if res.Error == nil {
		if res.RowsAffected > 0 {
			return true
		}
	}

	return false
}

func CreateTag(tag *Tag) error {
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()

	fmt.Println("tag =>", tag)
	fmt.Println("tag =>", tag.Name)
	fmt.Println("tag =>", tag.State)
	fmt.Println("createdId =>", tag.CreatedId)

	if err := DB.Model(&Tag{}).Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTag(tag *Tag) (effect int64, err error) {
	tag.UpdatedAt = time.Now()

	fmt.Println(tag)

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
