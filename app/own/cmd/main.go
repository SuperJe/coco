package main

import (
	"github.com/SuperJe/coco/app/own/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.RegisterAPI(r)
	if err := r.Run(":9090"); err != nil {
		panic(err)
	}
}
