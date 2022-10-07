package open

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"pd-go-server/models"
	e "pd-go-server/pkg/error"
	"pd-go-server/pkg/util"
)

func CreateUser(c *gin.Context) {
	var body map[string]interface{}

	c.BindJSON(&body)

	valid := validation.Validation{}

	fmt.Println(body["password"])

	valid.Required(body["account"], "account").Message("账户名不能为空")
	valid.Required(body["password"], "password").Message("密码不能为空")

	if valid.HasErrors() {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg": valid.Errors,
		})

		return
	}

	var user models.User

	user.Account = body["account"].(string)
	user.Password = util.MD5Cypto(body["password"].(string))

	mobile, hasMobile := body["mobile"].(string)
	nickName, hasNickName := body["nickname"].(string)
	email, hasEmail := body["email"].(string)

	if hasMobile {
		user.Mobile = mobile
	} else {
		user.Mobile = ""
	}

	if hasNickName {
		user.NickName = nickName
	} else {
		user.NickName = fmt.Sprintf("pd_user_%s", uuid.New().String())
	}

	if hasEmail {
		user.Email = email
	}

	err := models.CreateUser(&user)

	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"msg": e.GetMsg(e.SUCCESS),
	})
}
