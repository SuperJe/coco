package service

import (
	"fmt"
	"net/http"

	"github.com/SuperJe/coco/pkg/mongo/entity"
	"github.com/SuperJe/coco/pkg/util"
	"github.com/gin-gonic/gin"
)

func (s *Service) UserProgression(c *gin.Context) {
	req := &entity.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		fmt.Println("c.ShouldBindJSON err:", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println("receive req:", util.JSONString(req))
	c.AbortWithStatusJSON(http.StatusOK, "ok")
}
