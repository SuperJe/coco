package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastLevel string             `bson:"lastLevel"`
	Earned    *EarnedRewards     `bson:"earned" json:"earned,omitempty"`
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
