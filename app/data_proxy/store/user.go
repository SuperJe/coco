package store

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/pkg/common"
	"github.com/SuperJe/coco/pkg/util"
)

// User 对应admin库的sys_user表
type User struct {
	ID          int64     `xorm:"user_id"`
	TeacherID   int64     `xorm:"teacher_id"`
	Name        string    `xorm:"username"`
	TeacherName string    `xorm:"teacher_name"`
	Password    string    `xorm:"password"`
	Nick        string    `xorm:"nick_name"`
	Phone       string    `xorm:"phone"`
	RoleID      int32     `xorm:"role_id"`
	Sex         string    `xorm:"sex"`
	DeptID      int32     `xorm:"dept_id"`
	Remark      string    `xorm:"remark"`
	Status      int32     `xorm:"status"`
	CreateBy    int32     `xorm:"create_by"`
	UpdateBy    int32     `xorm:"update_by"`
	CreateAt    time.Time `xorm:"created_at"`
	UpdatedAt   time.Time `xorm:"updated_at"`
}

func (u User) TableName() string {
	return "admin.sys_user"
}

// Student 学生的基本信息
type Student struct {
	Name        string
	Password    string
	Phone       string
	Sex         string
	Class       string
	TeacherName string
}

// Valid 参数校验
func (s *Student) Valid() error {
	if s == nil {
		return fmt.Errorf("student nil")
	}
	if util.EmptyS(s.Class) || util.EmptyS(s.TeacherName) {
		return fmt.Errorf("invalid class or teacher name")
	}
	if err := common.CheckRegisterParam(s.Name, s.Password); err != nil {
		return errors.Wrap(err, "CheckRegisterParam err")
	}
	// 男0女1
	if s.Sex != strconv.FormatInt(common.GenderMale, 10) &&
		s.Sex != strconv.FormatInt(common.GenderFemale, 10) {
		return fmt.Errorf("invalid gender: %s", s.Sex)
	}
	const phoneLen = 11
	if len(s.Phone) != phoneLen {
		return fmt.Errorf("invalid phone number:%s", s.Phone)
	}
	return nil
}

// ToReq 转成注册请求结构
func (s *Student) ToReq() *model.RegisterReq {
	if s == nil {
		return nil
	}
	return &model.RegisterReq{
		Name:        s.Name,
		Pwd:         s.Password,
		Phone:       s.Phone,
		Sex:         s.Sex,
		Class:       s.Class,
		TeacherName: s.TeacherName,
	}
}

// StudentFromRegister 注册请求转换成student结构
func StudentFromRegister(r *model.RegisterReq) *Student {
	if r == nil {
		return nil
	}
	return &Student{
		Name:        r.Name,
		Password:    r.Pwd,
		Phone:       r.Phone,
		Sex:         r.Sex,
		Class:       r.Class,
		TeacherName: r.TeacherName,
	}
}

// NewStudent 新建一个学生账号
func (s *Store) NewStudent(ctx context.Context, student *Student) error {
	// 找到班级id
	sql := fmt.Sprintf("SELECT dept_id FROM admin.sys_dept WHERE dept_name = '%s'", student.Class)
	cID, err := s.queryID(ctx, sql, "dept_id")
	if err != nil {
		return err
	}
	// 找到教师id
	sql = fmt.Sprintf("SELECT teacher_id FROM admin.sys_user WHERE teacher_name = '%s'", student.TeacherName)
	tID, err := s.queryID(ctx, sql, "teacher_id")
	if err != nil {
		return err
	}
	// 密码hash
	hashed, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &User{
		TeacherID:   int64(tID),
		Name:        student.Name,
		TeacherName: student.TeacherName,
		Password:    string(hashed),
		Nick:        student.Name,
		Phone:       student.Phone,
		RoleID:      int32(common.RoleIDStudent),
		Sex:         student.Sex,
		DeptID:      int32(cID),
		Remark:      "Welcome!",
		Status:      common.StatusNormal,
		CreateBy:    1,
		UpdateBy:    1,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err = s.mysql.InsertOne(user)
	return err
}

func (s *Store) queryID(ctx context.Context, sql, field string) (int, error) {
	results, err := s.mysql.Context(ctx).QueryString(sql)
	if err != nil {
		// 没有update, 无需rollback
		return 0, err
	}
	if len(results) == 0 {
		return 0, fmt.Errorf("cannot find record")
	}
	return cast.ToInt(results[0][field]), nil
}
