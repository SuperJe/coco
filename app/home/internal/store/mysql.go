package store

// Reserve 预约表
type Reserve struct {
	ID       int64  `xorm:"id"`
	Name     string `xorm:"name"`
	Phone    string `xorm:"phone"`
	Grade    string `xorm:"grade"`
	Course   string `xorm:"course"`
	Period   string `xorm:"period"`
	Location string `xorm:"location"`
	Msg      string `xorm:"msg"`
	Ticket   string `xorm:"ticket"`
}

func (r *Reserve) TableName() string {
	return "reserve"
}
