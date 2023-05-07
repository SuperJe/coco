package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/pkg/common"
)

func (s *Service) BatchGetUserProgression(c *gin.Context) {
	req := &model.BatchGetUserProgressionReq{}
	rsp := &model.BatchGetUserProgressionRsp{}
	if err := c.BindQuery(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bind err")
		return
	}
	if err := json.Unmarshal([]byte(req.Bytes), &req.Names); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "unmarshal err")
		return
	}
	fmt.Printf("bs:%s, names:%+v\n\n", string(req.Bytes), req.Names)
	progressions, err := s.store.BatchGetUserProgressions(req.Names)
	if err != nil {
		rsp.Msg = "get err:" + err.Error()
		rsp.Code = common.ErrCodeDB
		c.AbortWithStatusJSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "success"
	rsp.CampProgressions = progressions.GroupByName()
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}
