package middleware

import "github.com/gin-gonic/gin"

func JSONMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Next()
	}
}
