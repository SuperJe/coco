package entity

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SuperJe/coco/pkg/util/encode"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name,omitempty" json:"name,omitempty"`
	NameLower         string             `bson:"nameLower,omitempty" json:"-"`
	LastLevel         string             `bson:"lastLevel,omitempty"`
	Earned            *EarnedRewards     `bson:"earned,omitempty" json:"earned,omitempty"`
	CreateOnHost      string             `bson:"createOnHost,omitempty" json:"-"`
	LastIP            string             `bson:"lastIP,omitempty" json:"-"`
	PreferredLanguage string             `bson:"preferredLanguage,omitempty" json:"-"`
	TestGroupNumber   int32              `bson:"testGroupNumber,omitempty" json:"-"`
	Anonymous         bool               `bson:"anonymous,omitempty" json:"-"`
	DateCreated       time.Time          `bson:"dateCreated,omitempty" json:"-"`
	V                 int32              `bson:"__v" json:"-"`
	Referrer          string             `bson:"referrer,omitempty" json:"-"`
	Birthday          string             `bson:"birthday,omitempty" json:"-"`
	Email             string             `bson:"email,omitempty" json:"-"`
	EmailLower        string             `bson:"emailLower,omitempty" json:"-"`
	PWDHash           string             `bson:"passwordHash,omitempty" json:"-"`
	Slug              string             `bson:"slug,omitempty" json:"-"`
	Emails            struct {
		GeneralNews struct {
			Enabled bool `bson:"enabled"`
		} `bson:"generalNews"`
	} `bson:"emails"`
}

func userNameSlugify(name string) string {
	// 将所有字母字符转换为小写
	name = strings.ToLower(name)
	// 使用空格分隔单词，并将单词连接成 slug 形式
	words := strings.Fields(name)
	return strings.Join(words, "-")
}

// NewUser 返回注册一个用户的所有信息
func NewUser(name, pwd string) *User {
	email := fmt.Sprintf("%d@qq.com", time.Now().UnixMilli())
	return &User{
		ID:                primitive.NewObjectID(),
		Name:              name,
		NameLower:         strings.ToLower(name),
		CreateOnHost:      "localhost:3020",
		LastIP:            "::ffff:172.17.0.1",
		PreferredLanguage: "zh-HANS",
		TestGroupNumber:   250,
		Anonymous:         false,
		DateCreated:       time.Now(),
		V:                 0,
		Referrer:          "http://localhost:3020/",
		Birthday:          time.Now().Format("2006-01-02 15:04:05"),
		Email:             email,
		EmailLower:        strings.ToLower(email),
		PWDHash:           encode.Sha512WithSalt(pwd, "pepper"),
		Slug:              userNameSlugify(name),
		Emails: struct {
			GeneralNews struct {
				Enabled bool `bson:"enabled"`
			} `bson:"generalNews"`
		}{
			GeneralNews: struct {
				Enabled bool `bson:"enabled"`
			}{
				Enabled: false,
			},
		},
	}
}

// Clone 深拷贝
func (u *User) Clone() *User {
	if u == nil {
		return nil
	}
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		LastLevel: u.LastLevel,
		Earned:    u.Earned,
	}
}

// CompletedLevelCount 已完成的关卡数
func (u *User) CompletedLevelCount() int32 {
	if u == nil || u.Earned == nil {
		return 0
	}
	return int32(len(u.Earned.Levels))
}

func (u *User) GemCount() int32 {
	if u == nil || u.Earned == nil {
		return 0
	}
	return int32(u.Earned.Gems)
}
