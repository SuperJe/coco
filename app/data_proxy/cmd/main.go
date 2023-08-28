package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SuperJe/coco/app/data_proxy/service"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	svc, err := service.NewService()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("../template/*")
	registerHandler(r, svc)
	if err := r.Run("0.0.0.0:7777"); err != nil {
		panic(err)
	}
}

// registerHandler 注册处理方法
func registerHandler(r *gin.Engine, svc *service.Service) {
	r.POST("/user_progression", svc.UpdateUserProgression)
	r.GET("/user_progression", svc.GetUserProgression)
	r.GET("/batch_user_progression", svc.BatchGetUserProgression)
	r.POST("/register", svc.Register)
	r.GET("/compiler", func(c *gin.Context) { c.HTML(http.StatusOK, "compiler.html", nil) })
	r.POST("/compiler/run", svc.CompilerRun)
}
