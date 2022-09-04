package v1

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"pd-go-server/models"
	"pd-go-server/pkg/error"
	"pd-go-server/pkg/util"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	pageInfo := util.FetchPage(c)

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(pageInfo.Offset, pageInfo.Limit, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(code, gin.H{
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	var tag models.Tag
	err := c.BindJSON(&tag)

	if err != nil {
		fmt.Println(err)

		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg": e.GetMsg(e.INVALID_PARAMS),
		})

		return
	}

	valid := validation.Validation{}

	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("最大程度不允许超过100字符")
	valid.Required(tag.CreatedId, "author").Message("必须传入创建人信息")
	valid.Min(tag.CreatedId, 1, "author").Message("无效创建人")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	if !valid.HasErrors() {
		if models.ExistTagByName(tag.Name) {
			c.JSON(e.ERROR_EXIST_TAG, gin.H{
				"msg": e.GetMsg(e.ERROR_NOT_EXIST_TAG),
			})

			return
		} else {
			c.JSON(e.INVALID_PARAMS, gin.H{
				"msg": e.GetMsg(e.INVALID_PARAMS),
			})

			return
		}
	}

	code := e.SUCCESS
	createError := models.CreateTag(&tag)

	if createError != nil {
		code = e.INTERNAL_ERROR
	}

	c.JSON(code, gin.H{
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func UpdateTag(c *gin.Context) {
	var tag models.Tag

	c.BindJSON(&tag)
	id := util.StringToInt64(c.Param("id"), 0)

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("更新标签ID不能为空")

	if !valid.HasErrors() {
		if !models.ExistTagById(id) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  e.GetMsg(e.ERROR_NOT_EXIST_TAG),
				"code": e.GetMsg(e.ERROR_NOT_EXIST_TAG),
			})

			return
		}

		tag.ID = uint(id)
		rows, err := models.UpdateTag(&tag)

		if err != nil {
			c.JSON(e.INTERNAL_ERROR, gin.H{
				"msg": e.GetMsg(e.INTERNAL_ERROR),
			})

			return
		}

		data := make(map[string]interface{})

		data["effectRows"] = rows

		c.JSON(e.SUCCESS, gin.H{
			"msg":  e.GetMsg(e.SUCCESS),
			"data": data,
		})

	}

}

func DeleteTag(c *gin.Context) {
	id := util.StringToInt64(c.Param("id"), 0)

	if id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "无效的删除id",
		})

		return
	}

	rows, err := models.DeleteTag(id)

	if err != nil {
		fmt.Println("删除标签异常", err)
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"code": e.INTERNAL_ERROR,
			"msg":  e.GetMsg(e.INTERNAL_ERROR),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rows,
	})
}
