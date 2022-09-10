package models

import (
	"gorm.io/gorm"
)

type Article struct {
	Model
	Title     string `gorm:"title" json:"title"`
	Desc      string `gorm:"desc" json:"desc"`
	Content   string `gorm:"content" json:"content"`
	State     uint   `gorm:"state" json:"state"`
	CreatedId uint   `gorm:"created_id" json:"created_id"`
}

type ArticleWithTags struct {
	Article
	Tags []Tag `json:"tags"`
}

func ExistArticleByID(id int64) bool {
	var article Article
	DB.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(query map[string]interface{}) (count *int64) {
	DB.Model(&Article{}).Where(query).Count(count)

	return
}

func GetArticles(limit int, offset int, query map[string]interface{}) (data []ArticleWithTags, err error) {
	err = nil

	var articles []Article

	needTag, isOk := query["tag"]
	orderStr, isEffectOrder := query["order"]

	delete(query, "tag")
	delete(query, "order")

	if !isEffectOrder {
		orderStr = "id desc"
	}

	query["State"] = 1

	res := DB.Model(&Article{}).Offset(offset).Limit(limit).Where(query).Omit("Content", "State", "UpdatedAt", "DeletedAt").Order(orderStr).Find(&articles)

	if err = res.Error; err != nil {
		return
	}

	if isOk && needTag != nil {
		for _, article := range articles {
			tags, e := GetTagsByArticleId(10, 0, article.ID)

			if err = e; err != nil {
				return
			}
			data = append(data, ArticleWithTags{
				article,
				*tags,
			})
		}
	}
	return
}

func GetArticleById(id uint) (data ArticleWithTags, err error) {
	err = nil
	var article Article
	res := DB.Model(Article{}).Where(map[string]interface{}{"ID": id}).First(&article)

	if err = res.Error; err != nil {
		return
	}

	tags, err := GetTagsByArticleId(100, 0, id)

	if err != nil {
		return
	}

	data.Article = article
	data.Tags = *tags

	return
}

func AddArticle(article *Article, tags *[]Tag) (err error) {
	err = nil

	DB.Transaction(func(tx *gorm.DB) error {
		UseTransaction(tx)

		if err = tx.Create(article).Error; err != nil {
			return err
		}

		for i, _ := range *tags {
			id := (*tags)[i].ID

			if exist := ExistTagById(int64(id)); exist {
				continue
			}

			if exist := ExistTagByName(&(*tags)[i]); !exist {
				if err := CreateTag(&(*tags)[i]); err != nil {
					return err
				}
			}
		}

		if _, err := CreateTagByArticle(article, tags); err != nil {
			return err
		}

		return nil
	})

	TransactionEnd()

	return
}

func EditArticle() {

}

func DeleteArticle() {

}
