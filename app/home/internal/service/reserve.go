package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	"github.com/SuperJe/coco/app/home/internal/common"
	db "github.com/SuperJe/coco/pkg/mysql"
	"github.com/SuperJe/coco/pkg/util"
)

func (s *Service) Reserve(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &common.ReserveReq{}
	rsp := &common.ReserveRsp{}
	if err := c.ShouldBind(req); err != nil {
		fmt.Println("c.ShouldBindJSON err:", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !req.IsValid() {
		fmt.Println("req invalid:", req)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(req)
	ticket, err := s.addReservation(ctx, req)
	if err != nil {
		rsp.Code = -1
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			rsp.Code = int32(errMySQL.Number)
			fmt.Println("db err:", err.Error())
		}
		rsp.Msg = err.Error()
		c.AbortWithStatusJSON(http.StatusOK, rsp)
		return
	}
	rsp.Ticket = ticket
	c.AbortWithStatusJSON(http.StatusOK, rsp)
	return
}

func (s *Service) addReservation(ctx context.Context, req *common.ReserveReq) (string, error) {
	ticket := util.RandString(10)
	reserve := req.ToReserve(ticket)
	if err := s.store.AddReservation(ctx, reserve); err != nil {
		if db.IsDupErr(err) {
			r, err := s.store.GetReservation(ctx, reserve.Phone)
			if err != nil {
				return "", err
			}
			return r.Ticket, nil
		}
		return "", err
	}
	return ticket, nil
}
