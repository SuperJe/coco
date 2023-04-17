package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Level 保存在level集合中的关卡具体结构
type Level struct {
	Slug         string             `bson:"slug"`
	EnName       string             `bson:"name"`
	CampaignName string             `bson:"campaign"` // 对应Campaign集合的name
	I18Ns        map[string]*I18N   `bson:"i18n"`
	Original     primitive.ObjectID `bson:"original,omitempty"`
}
