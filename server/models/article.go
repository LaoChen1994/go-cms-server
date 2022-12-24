package models

import (
	"fmt"
	"gorm.io/gorm"
	"pd-go-server/pkg/util"
)

type Article struct {
	Model
	Title     string `gorm:"title" json:"title"`
	Desc      string `gorm:"desc" json:"desc"`
	Content   string `gorm:"content" json:"content"`
	State     uint   `gorm:"state" json:"state"`
	CreatedId uint   `gorm:"created_id" json:"createdId"`
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

	res := DB.Model(&Article{}).Offset(offset).Limit(limit).Where(query).Omit("Content", "UpdatedAt", "DeletedAt").Order(orderStr).Find(&articles)

	fmt.Println(res)

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

	// @todo 后续考虑标签过多的场景
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

func EditArticle(article *Article, tags *[]Tag) (err error) {
	err = nil

	DB.Transaction(func(tx *gorm.DB) error {
		UseTransaction(tx)

		fmt.Println("ID =>", article.ID)

		err := tx.Model(&Article{}).Where("id = ?", article.ID).Omit("ID").Updates(&article).Error

		if err != nil {
			return err
		}

		articleCount := GetTagsCountByArticleId(article.ID)
		tagsLen := len(*tags)

		if articleCount == 0 && tagsLen == 0 {
			return nil
		}

		relations, queryTagError := GetContentTagsByArticleId(100, 0, article.ID)

		if queryTagError != nil {
			err = queryTagError
			return queryTagError
		}

		var newRelation []Tag

		newTags := util.Filter(*tags, func(tag Tag, i int) bool {
			if tag.ID > 0 {
				match := util.Find(relations, func(contentTag ContentTag, i int) bool {
					return uint(contentTag.TagId) == tag.ID
				})

				if match.ID <= 0 {
					// 说明标签存在但是关联关系不在
					newRelation = append(newRelation, tag)
				}
			}

			return tag.ID == 0
		})

		deleteTag := util.Filter(relations, func(relation ContentTag, i int) bool {
			matchTag := util.Find(*tags, func(tag Tag, i int) bool {
				return tag.ID == uint(relation.TagId)
			})

			return matchTag.ID <= 0
		})

		if len(newTags) > 0 {
			for i, _ := range newTags {
				if exist := ExistTagByName(&newTags[i]); !exist {
					if err := CreateTag(&newTags[i]); err != nil {
						return err
					}
				}
			}

			if _, err := CreateTagByArticle(article, &newTags); err != nil {
				return err
			}
		}

		if len(newRelation) > 0 {
			if _, err := CreateTagByArticle(article, &newRelation); err != nil {
				return err
			}
		}

		if len(deleteTag) > 0 {
			if err := DeleteContentTagByIds(&deleteTag); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil
	})

	TransactionEnd()
	return
}

func DeleteArticle(id uint) (err error) {
	var article Article

	err = nil

	if err = DB.Model(&Article{}).Where("ID = ", id).Delete(&article).Error; err != nil {
		return err
	}

	return
}
