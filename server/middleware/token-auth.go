package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pd-go-server/models"
	e "pd-go-server/pkg/error"
	"pd-go-server/pkg/jwt"
	"strings"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.String(), "/api/open") {
			c.Next()
			return
		}

		token, err := c.Cookie("auth_token")

		if err != nil {
			c.JSON(e.INVALID_PARAMS, gin.H{
				"data": e.ERROR_AUTH,
				"msg":  e.GetMsg(e.ERROR_AUTH),
			})

			c.Abort()

			return
		}

		claims, jwtError := jwt.ParseToken(token)

		if jwtError != nil {
			fmt.Println(jwtError.Error())
			c.JSON(e.INVALID_PARAMS, gin.H{
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

		isValid := models.IsValidUser(user)

		if !isValid {
			c.JSON(e.INVALID_PARAMS, gin.H{
				"data": e.ERROR_AUTH,
				"msg":  e.GetMsg(e.ERROR_AUTH),
			})
			c.Abort()

			return
		}

		c.Set("userId", user.ID)
		c.Next()
	}
}
