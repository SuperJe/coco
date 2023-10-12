package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"

	"github.com/SuperJe/coco/app/home/internal/service"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	svc, err := service.NewService()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Static("/static", "../static")
	r.LoadHTMLGlob("../template/*")
	registerHandler(r, svc)
	if err := r.Run("0.0.0.0:7471"); err != nil {
		panic(err)
	}
}

// registerHandler 注册处理方法
func registerHandler(r *gin.Engine, svc *service.Service) {
	r.GET("", svc.GetIndex)
	r.GET("/generic", func(c *gin.Context) {
		c.HTML(http.StatusOK, "generic.html", nil)
	})
	r.GET("/elements", func(c *gin.Context) {
		c.HTML(http.StatusOK, "elements.html", nil)
	})

	r.POST("/reserve", svc.Reserve)
}
