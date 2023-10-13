package common

import (
	"github.com/SuperJe/coco/app/home/internal/store"
	"github.com/SuperJe/coco/pkg/util"
)

var GradeMap = map[string]string{
	"secondary_low":  "小学1-3年级",
	"secondary_high": "小学4-6年级",
	"junior":         "初中",
	"high":           "高中",
}

var PeriodMap = map[string]string{
	"t1": "周六09:30-11:30",
	"t2": "周六14:00-16:00",
	"t3": "周六16:30-18:30",
	"t4": "周六19:00-21:00",
	"t5": "周日09:30-11:30",
	"t6": "周日14:00-16:00",
	"t7": "周日16:30-18:30",
	"t8": "周日19: 00 - 21:00",
}

var LocationMap = map[string]string{
	"gudishi":   "古地石校区",
	"bainaohui": "百脑汇校区",
}

type BaseRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type ReserveReq struct {
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Grade    string `form:"grade"`
	Course   string `form:"course"`
	Period   string `form:"period"`
	Location string `form:"location"`
	Msg      string `form:"message"`
}

func (rr *ReserveReq) IsValid() bool {
	if rr == nil {
		return false
	}
	return !util.ExistEmptyString(rr.Name, rr.Phone, rr.Grade, rr.Course, rr.Period, rr.Location)
}

func (rr *ReserveReq) ToReserve(ticket string) *store.Reserve {
	if rr == nil {
		return nil
	}
	return &store.Reserve{
		Name:     rr.Name,
		Phone:    rr.Phone,
		Grade:    GradeMap[rr.Grade],
		Course:   rr.Course,
		Period:   PeriodMap[rr.Period],
		Location: LocationMap[rr.Location],
		Msg:      rr.Msg,
		Ticket:   ticket,
	}
}

type ReserveRsp struct {
	BaseRsp
	Ticket string `json:"ticket"`
}
