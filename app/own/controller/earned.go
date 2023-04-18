package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func earned(c *gin.Context) {
	res := struct {
		Msg string `json:"msg"`
	}{"ok"}
	c.AbortWithStatusJSON(http.StatusOK, res)
}
