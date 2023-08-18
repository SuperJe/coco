package model

type BaseRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type UpdateUserProgressionReq struct {
	Name            string               `json:"name"`
	Completed       int32                `json:"completed"`
	Gems            int32                `json:"gems"`
	LastLevel       string               `json:"last_level"`
	CampProgression *CampaignProgression `json:"camp_progression"`
}

type UpdateUserProgressionRsp struct {
	BaseRsp
}

type GetUserProgressionReq struct {
	Name string `form:"name" json:"name"`
}

type GetUserProgressionRsp struct {
	BaseRsp
	Completed           int32                `json:"completed"`
	Gems                int32                `json:"gems"`
	LastLevel           string               `json:"last_level"`
	CampaignProgression *CampaignProgression `json:"camp_progression"`
}

type BatchGetUserProgressionReq struct {
	Bytes string   `form:"names" json:"names"`
	Names []string `json:"-"`
}

type BatchGetUserProgressionRsp struct {
	BaseRsp
	CampProgressions map[string]*CampaignProgression `json:"camp_progressions"`
}

// RegisterReq register request body
type RegisterReq struct {
	Name        string `json:"name"`
	Pwd         string `json:"pwd"`
	Phone       string `json:"phone"`
	Sex         string `json:"sex"`
	Class       string `json:"class"`
	TeacherName string `json:"teacher_name"`
}

type RegisterRsp struct {
	BaseRsp
}
