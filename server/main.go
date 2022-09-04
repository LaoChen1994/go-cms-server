package main

import (
	"fmt"
	"net/http"
	"pd-go-server/pkg/setting"
	"pd-go-server/routers"
)

func main() {
	// 创建一个Gin的路由中间件
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server is work on ", setting.HttpPort)

	server.ListenAndServe()
}
