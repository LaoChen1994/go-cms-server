package models

import (
	"fmt"
	"pd-go-server/pkg/setting"
)

type ContentTag struct {
	Model
	ContentId int64 `gorm:"column:content_id" json:"contentId" gorm:"index"`
	TagId     int64 `gorm:"column:tag_id" json:"tagId" gorm:"index"`
	CreatedId uint  `gorm:"column:created_id" json:"createdId"`
}

func (*ContentTag) TableName() string {
	return fmt.Sprintf("%s%s", setting.DatabaseConf.Prefix, "article_tag")
}

func CreateTagByArticle(article *Article, tags *[]Tag) (effectRows int64, err error) {
	db := GetDB()

	var contentTags []ContentTag

	effectRows = 0
	err = nil

	articleId := article.ID
	createdId := article.CreatedId

	for _, tag := range *tags {
		contentTags = append(contentTags, ContentTag{
			ContentId: int64(articleId),
			TagId:     int64(tag.ID),
			CreatedId: createdId,
		})
	}

	if createErr := db.Model(&ContentTag{}).Omit("CreatedAt", "UpdatedAt").Create(&contentTags).Error; createErr != nil {
		err = createErr
		return
	}

	return
}

func GetTagsByArticleId(limit int, offset int, articleId uint) (tags *[]Tag, err error) {
	err = nil

	res := DB.Model(&ContentTag{}).
		Limit(limit).Offset(offset).
		Where("content_id = ?", articleId).
		Select("`blog_tag`.name, `blog_tag`.id, `blog_tag`.parent_id, `blog_tag`.state").
		Joins("left join `blog_tag` on `blog_tag`.id = `blog_article_tag`.tag_id").
		Where("`blog_tag`.state = 1").
		Find(&tags)

	if err = res.Error; err != nil {
		return
	}

	return
}

func GetContentTagsByArticleId(limit int, offset int, articleId uint) (data []ContentTag, err error) {
	err = DB.Model(&ContentTag{}).Limit(limit).Offset(offset).Where("content_id = ?", articleId).Find(&data).Error
	return
}

func GetTagsCountByArticleId(articleId uint) (count int64) {

	DB.Model(&ContentTag{}).Where("content_id = ?", articleId).Count(&count)

	return
}

func DeleteContentTagByIds(contentTag *[]ContentTag) error {
	return DB.Model(&ContentTag{}).Delete(contentTag).Error
}

func GetArticlesByTag(limit int, offset int, tagId uint) (articles *[]ContentTag, err error) {
	err = nil

	res := DB.Limit(limit).Offset(offset).Where("TagId = ?", tagId).Find(&articles)

	if err = res.Error; err != nil {
		return
	}

	return
}
