package model

// UserProgressionReq user_progression接口body
type UserProgressionReq struct {
	Name        string `json:"name"`
	LastLevel   string `json:"last_level"`
	CompetedNum int64  `json:"competed_num"`
	TotalLevels int64  `json:"total_levels"`
}

// UserProgressionRsp user_progression接口response
type UserProgressionRsp struct {
}
