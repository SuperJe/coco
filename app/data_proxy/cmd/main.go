package main

import (
	"github.com/SuperJe/coco/app/data_proxy/service"
	"github.com/gin-gonic/gin"
)

func main() {
	svc, err := service.NewService()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	registerHandler(r, svc)
	if err := r.Run("127.0.0.1:7777"); err != nil {
		panic(err)
	}
}

// registerHandler 注册处理方法
func registerHandler(r *gin.Engine, svc *service.Service) {
	r.POST("/user_progression", svc.UpdateUserProgression)
	r.GET("/user_progression", svc.GetUserProgression)
}
