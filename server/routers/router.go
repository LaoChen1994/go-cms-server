package routers

import (
	"github.com/gin-gonic/gin"
	"pd-go-server/middleware"
	"pd-go-server/pkg/setting"
	"pd-go-server/routers/api/open"
	"pd-go-server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.TokenAuthMiddleware())
	r.Use(middleware.JSONMiddleware())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/cms")
	apiOpen := r.Group("/api/open")

	apiv1.GET("/tags", v1.GetTags)
	apiv1.POST("/tags", v1.AddTag)
	apiv1.PUT("/tags/:id", v1.UpdateTag)
	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	apiv1.GET("/articles", v1.GetArticles)
	apiv1.GET("/article/:id", v1.GetArticle)
	apiv1.POST("/article", v1.AddArticle)
	apiv1.PUT("/article", v1.UpdateArticle)
	apiv1.DELETE("/article/:id", v1.DeleteArticle)

	// 登录
	apiOpen.POST("/user/login", open.RequestLogin)
	// 创建token
	apiOpen.POST("/user", open.CreateUser)
	// 创建用户token
	apiOpen.GET("/user/auth/:id", open.GetAuth)
	// 确认用户登录并返回相关信息
	apiOpen.GET("/user/auth", open.AuthByToken)
	// 用户登出
	apiOpen.POST("/user/loginout", open.RequestLoginout)

	return r
}
