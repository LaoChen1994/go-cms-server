package open

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"pd-go-server/models"
	e "pd-go-server/pkg/error"
	"pd-go-server/pkg/jwt"
	"pd-go-server/pkg/setting"
	"pd-go-server/pkg/util"
	"strconv"
)

func GetAuth(c *gin.Context) {
	if setting.RunMode != "debug" {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"msg": "生产环境不可用",
		})

		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg": "无效用户ID",
		})

		return
	}

	user, _ := models.GetUserById(uint(id))

	fmt.Println(user.Password)

	token, tokenError := jwt.GenerateToken(user.Account, user.Password)

	if tokenError != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg": tokenError.Error(),
		})

		return
	}

	c.JSON(e.SUCCESS, gin.H{
		"msg": e.GetMsg(e.SUCCESS),
		"data": map[string]interface{}{
			"token": token,
		},
	})

}

func RequestLogin(c *gin.Context) {
	var body map[string]interface{}

	c.BindJSON(&body)

	valid := validation.Validation{}

	valid.Required(body["account"], "account").Message("用户名不能为空")
	valid.Required(body["password"], "password").Message("密码不能为空")

	passwd := util.MD5Cypto(body["password"].(string))
	account := body["account"].(string)

	isValid := models.IsValidUser(&models.User{
		Password: passwd,
		Account:  account,
	})

	if !isValid {
		c.JSON(e.ERROR_USER, gin.H{
			"msg": e.GetMsg(e.ERROR_USER),
		})
		return
	}

	token, err := jwt.GenerateToken(account, passwd)

	if err != nil {
		c.JSON(e.ERROR_AUTH_TOKEN, gin.H{
			"msg": e.GetMsg(e.ERROR_AUTH_TOKEN),
		})

		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("auth_token", token, 7*3600*24, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"msg": e.GetMsg(e.SUCCESS),
	})
}

func AuthByToken(c *gin.Context) {
	token, err := c.Cookie("auth_token")

	if err != nil {
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"msg":  "无法获取token",
			"code": e.ERROR_AUTH_CHECK_TOKEN_FAIL,
		})

		c.Abort()
		return
	}

	claims, jwtError := jwt.ParseToken(token)

	if jwtError != nil {
		fmt.Println(jwtError.Error())
		c.JSON(e.INTERNAL_ERROR, gin.H{
			"data": e.ERROR_AUTH_TOKEN,
			"msg":  e.GetMsg(e.ERROR_AUTH_TOKEN),
		})

		c.Abort()
		return
	}

	user := models.User{
		Account:  claims.Account,
		Password: claims.Password,
	}

	isValid := models.IsValidUser(&user)

	if isValid {
		c.JSON(e.SUCCESS, gin.H{
			"code": e.SUCCESS,
			"data": map[string]interface{}{
				"name":   user.NickName,
				"id":     user.ID,
				"email":  user.Email,
				"mobile": user.Mobile,
			},
		})
	} else {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"code": e.ERROR_USER,
			"msg":  "无效用户",
		})
	}
}

func RequestLoginout(c *gin.Context) {
	_, err := c.Cookie("auth_token")
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "未登录",
		})
		c.Abort()

		return
	}

	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie(
		"auth_token",
		"", -1,
		"/",
		"",
		true,
		true,
	)

	c.JSON(200, gin.H{
		"msg": "已退出登录",
	})
	c.Abort()

}
