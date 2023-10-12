package common

const (
	// ErrCodeDB DB操作错误
	ErrCodeDB         = 10001
	ErrCodeDBNotExist = 10002
	ErrCodeDBExisted  = 10003

	// ErrCompile 编译失败
	ErrCompile = 20001
)

const (
	RoleID = iota
	RoleIDAdmin
	RoleIDStudent
	RoleIDTeacher

	RoleTeacher = "teacher"
	RoleStudent = "student"
	RoleAdmin   = "admin"
)

const (
	GenderMale = iota
	GenderFemale
)

const (
	StatusBanned = iota + 1
	StatusNormal
)
