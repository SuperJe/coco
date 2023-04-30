package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/app/data_proxy/store"
	"github.com/SuperJe/coco/pkg/common"
)

func (s *Service) UpdateUserProgression(c *gin.Context) {
	req := &model.UpdateUserProgressionReq{}
	rsp := &model.UpdateUserProgressionRsp{}
	if err := c.ShouldBindJSON(req); err != nil {
		fmt.Println("c.ShouldBindJSON err:", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := s.store.UpsertUserProgression(store.NewUserProgression(req)); err != nil {
		rsp.Msg = "insert err:" + err.Error()
		rsp.Code = common.ErrCodeDB
		c.AbortWithStatusJSON(http.StatusOK, rsp)
	}
	rsp.Msg = "success"
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}
