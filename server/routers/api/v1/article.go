package v1

import (
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"pd-go-server/models"
	e "pd-go-server/pkg/error"
	"pd-go-server/pkg/util"
	"strconv"
)

func GetArticles(c *gin.Context) {
	query := make(map[string]interface{})

	needTag := c.Query("tag")
	pageInfo := util.FetchPage(c)

	// 查询参数暂时只支持按标签进行查询
	if needTag == "true" {
		query["tag"] = true
	} else {
		query["tag"] = nil
	}

	articles, err := models.GetArticles(pageInfo.Limit, pageInfo.Offset, query)

	if err != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg":   e.GetMsg(e.INTERNAL_ERROR),
			"error": err,
		})

		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"msg":  e.GetMsg(e.SUCCESS),
		"data": articles,
	})
}

func GetArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": "无效的文章id",
		})

		return
	}

	article, articleErr := models.GetArticleById(uint(id))

	if articleErr != nil || article.ID <= 0 {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": "无对应的文章",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  e.GetMsg(e.SUCCESS),
		"data": article,
	})
}

func AddArticle(c *gin.Context) {
	var body = make(map[string]interface{})
	var tags []models.Tag

	var article models.Article

	c.BindJSON(&body)

	valid := validation.Validation{}

	valid.Required(body["title"], "title").Message("文章标题不能为空")
	valid.Required(body["desc"], "description").Message("文章描述不能为空")
	valid.Required(body["content"], "content").Message("文章内容不能为空")
	valid.Range(int(body["state"].(float64)), 0, 1, "state").Message("上下架状态应该为0或1")

	if valid.HasErrors() {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg":   e.GetMsg(e.INVALID_PARAMS),
			"error": valid.ErrorMap(),
		})

		return
	}

	// todo 等到用户系统建立起来以后操作
	createdId := 1

	article.Title = body["title"].(string)
	article.Content = body["content"].(string)
	article.Desc = body["desc"].(string)
	article.State = uint(body["state"].(float64))
	article.CreatedId = uint(createdId)

	tagsBody := body["tags"].([]interface{})

	for _, t := range tagsBody {
		tag := models.Tag{}

		tBody, _ := t.(map[string]interface{})

		if id, ok := tBody["id"].(float64); ok {
			if id > 0 {
				tag.ID = uint(id)
				tags = append(tags, tag)
				continue
			}
		}

		if name, ok := tBody["name"].(string); ok {
			tag.Name = name
		} else {
			continue
		}

		if parentId, ok := tBody["pid"].(float64); ok {
			tag.ParentId = int(parentId)
		}

		if state, ok := tBody["state"].(uint); ok {
			tag.State = state
		} else {
			tag.State = 0
		}

		tag.CreatedId = uint(createdId)

		tags = append(tags, tag)
	}

	if err := models.AddArticle(&article, &tags); err != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg": e.GetMsg(e.INTERNAL_ERROR),
		})
		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"msg": e.GetMsg(e.SUCCESS),
	})
}

func UpdateArticle(c *gin.Context) {
	var body = make(map[string]interface{})
	var tags []models.Tag

	var article models.Article

	c.BindJSON(&body)

	valid := validation.Validation{}

	valid.Required(body["id"], "id").Message("更新文章必须存在ID参数")
	valid.Required(body["title"], "title").Message("文章标题不能为空")
	valid.Required(body["desc"], "description").Message("文章描述不能为空")
	valid.Required(body["content"], "content").Message("文章内容不能为空")
	valid.Range(int(body["state"].(float64)), 0, 1, "state").Message("上下架状态应该为0或1")

	if valid.HasErrors() {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg":   e.GetMsg(e.INVALID_PARAMS),
			"error": valid.ErrorMap(),
		})

		return
	}

	// todo 等到用户系统建立起来以后操作
	createdId := 1

	article.ID = uint(body["id"].(float64))
	article.Title = body["title"].(string)
	article.Content = body["content"].(string)
	article.Desc = body["desc"].(string)
	article.State = uint(body["state"].(float64))
	article.CreatedId = uint(createdId)

	tagsBody := body["tags"].([]interface{})

	for _, t := range tagsBody {
		tag := models.Tag{}

		tBody, _ := t.(map[string]interface{})

		if id, ok := tBody["id"].(float64); ok {
			if id > 0 {
				tag.ID = uint(id)
				tags = append(tags, tag)
				continue
			}
		}

		if name, ok := tBody["name"].(string); ok {
			tag.Name = name
		} else {
			continue
		}

		if parentId, ok := tBody["pid"].(float64); ok {
			tag.ParentId = int(parentId)
		}

		if state, ok := tBody["state"].(uint); ok {
			tag.State = state
		} else {
			tag.State = 0
		}

		tag.CreatedId = uint(createdId)

		tags = append(tags, tag)
	}

	if err := models.EditArticle(&article, &tags); err != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg": e.GetMsg(e.INTERNAL_ERROR),
		})
		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"msg": e.GetMsg(e.SUCCESS),
	})
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil || id <= 0 {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg":  e.GetMsg(e.INTERNAL_ERROR),
			"data": map[string]interface{}{},
		})
		return
	}

	err = models.DeleteArticle(uint(id))

	if err != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg":  e.GetMsg(e.INTERNAL_ERROR),
			"data": map[string]interface{}{},
		})
		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"msg": e.GetMsg(e.SUCCESS),
	})
}
