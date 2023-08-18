package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/app/data_proxy/store"
	"github.com/SuperJe/coco/pkg/common"
)

func (s *Service) Register(c *gin.Context) {
	req := &model.RegisterReq{}
	rsp := &model.RegisterRsp{}
	if err := c.ShouldBindJSON(req); err != nil {
		fmt.Println("c.ShouldBindJSON err:", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	student := store.StudentFromRegister(req)
	if err := s.store.NewStudent(ctx, student); err != nil {
		rsp.Msg = "new student err:" + err.Error()
		rsp.Code = common.ErrCodeDB
		c.AbortWithStatusJSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "success"
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}
