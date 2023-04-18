package controller

import "github.com/gin-gonic/gin"

func RegisterAPI(r *gin.Engine) {
	r.GET("/earned", earned)
}
