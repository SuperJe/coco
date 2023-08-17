package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Level 保存在level集合中的关卡具体结构
type Level struct {
	Slug                 string             `bson:"slug"`
	CampaignName         string             `bson:"campaign"` // 对应Campaign集合的name
	EnName               string             `bson:"name"`
	Desc                 string             `bson:"description"`
	Type                 string             `bson:"type"`
	Kind                 string             `bson:"kind,omitempty"`
	Concepts             []string           `bson:"concepts"`
	ScoreTypes           []string           `bson:"scoreTypes,omitempty"`
	RestrictedProperties []string           `bson:"restrictedProperties,omitempty"`
	RequiredProperties   []string           `bson:"requiredProperties,omitempty"`
	I18Ns                map[string]*I18N   `bson:"i18n"`
	Original             primitive.ObjectID `bson:"original,omitempty"`
}

// ToBrief 将Level结构中有的部分复制到LevelBrief结构中
// 当前Position是写死的，PrimaryConcepts是没有的
func (l *Level) ToBrief() *LevelBrief {
	return &LevelBrief{
		Slug:                   l.Slug,
		Campaign:               l.CampaignName,
		Name:                   l.EnName,
		Desc:                   l.Desc,
		Type:                   l.Type,
		Kind:                   l.Kind,
		Concepts:               l.Concepts,
		ScoreTypes:             l.ScoreTypes,
		RestrictedProperties:   l.RestrictedProperties,
		RequiredProperties:     l.RequiredProperties,
		I18Ns:                  l.I18Ns,
		Position:               &LevelPos{X: 50.0, Y: 50.0},
		Original:               l.Original,
		HidesRealTimePlayback:  true,
		HidesCodeToolbar:       true,
		HidesSay:               true,
		HidesHUD:               true,
		HidesRunShortcut:       true,
		HidesPlayButton:        true,
		HidesSubmitUntilRun:    true,
		LockDefaultCode:        false,
		BackspaceThrottle:      true,
		DisableSpaces:          3,
		AutocompleteFontSizePx: 20,
		CampaignIndex:          0,
	}
}
