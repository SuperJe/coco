package service

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Register(c *gin.Context) {
	// req := &model.RegisterReq{}
	// rsp := &model.RegisterRsp{}
	// if err := c.ShouldBindJSON(req); err != nil {
	// 	fmt.Println("c.ShouldBindJSON err:", err.Error())
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }
	// if err := s.store.UpsertUserProgression(store.NewUserProgression(req)); err != nil {
	// 	rsp.Msg = "insert err:" + err.Error()
	// 	rsp.Code = common.ErrCodeDB
	// 	c.AbortWithStatusJSON(http.StatusOK, rsp)
	// }
	// rsp.Msg = "success"
	// c.AbortWithStatusJSON(http.StatusOK, rsp)
}
