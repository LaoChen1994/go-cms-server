package open

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pd-go-server/models"
	e "pd-go-server/pkg/error"
	"pd-go-server/pkg/jwt"
	"pd-go-server/pkg/setting"
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
