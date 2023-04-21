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
	if err := r.Run(":9090"); err != nil {
		panic(err)
	}
}

// registerHandler 注册处理方法
func registerHandler(r *gin.Engine, svc *service.Service) {
	r.POST("/user", svc.UserProgression)
}
