package common

const (
	// ErrCodeDB DB操作错误
	ErrCodeDB         = 10001
	ErrCodeDBNotExist = 10002
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
