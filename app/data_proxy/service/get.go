package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/pkg/common"
)

func (s *Service) GetUserProgression(c *gin.Context) {
	req := &model.GetUserProgressionReq{}
	rsp := &model.GetUserProgressionRsp{}
	if err := c.BindQuery(req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	up, err := s.store.GetUserProgression(req.Name)
	if err != nil {
		rsp.Msg = "get err:" + err.Error()
		rsp.Code = common.ErrCodeDB
		c.AbortWithStatusJSON(http.StatusOK, rsp)
		return
	}
	if up == nil {
		rsp.Msg = "user not exist"
		rsp.Code = common.ErrCodeDBNotExist
		c.AbortWithStatusJSON(http.StatusOK, rsp)
		return
	}
	rsp.Msg = "success"
	rsp.Gems = up.Gems
	rsp.Completed = up.Completed
	rsp.CampaignProgression = up.Detail
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}
