package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Campaign 战役, 如Dungeon, Desert等大战役, 其中的Levels字段是简要字段，并非全部字段
type Campaign struct {
	Slug     string                 `bson:"slug"`
	Name     string                 `bson:"name"`
	FullName string                 `bson:"fullName"`
	Levels   map[string]*LevelBrief `bson:"levels"`
	I18Ns    map[string]*I18N       `bson:"i18n"`
}

// LevelBrief 在战役中有非常多的小关卡, 此为保存在Campaign结构中的小关卡的关键数据结构
type LevelBrief struct {
	Slug                   string             `bson:"slug"`
	Campaign               string             `bson:"campaign"`
	Name                   string             `bson:"name"`
	Desc                   string             `bson:"description"`
	Type                   string             `bson:"type"`
	Kind                   string             `bson:"kind,omitempty"`
	HidesRealTimePlayback  bool               `bson:"hidesRealTimePlayback"`
	HidesCodeToolbar       bool               `bson:"hidesCodeToolbar"`
	HidesSay               bool               `bson:"hidesSay"`
	HidesHUD               bool               `bson:"hidesHUD"`
	HidesRunShortcut       bool               `bson:"hidesRunShortcut"`
	HidesPlayButton        bool               `bson:"hidesPlayButton"`
	HidesSubmitUntilRun    bool               `bson:"hidesSubmitUntilRun"`
	LockDefaultCode        bool               `bson:"lockDefaultCode"`
	BackspaceThrottle      bool               `bson:"backspaceThrottle"`
	DisableSpaces          int                `bson:"disableSpaces"`
	AutocompleteFontSizePx int                `bson:"autocompleteFontSizePx"`
	CampaignIndex          int                `bson:"campaignIndex"`
	Concepts               []string           `bson:"concepts"`
	ScoreTypes             interface{}        `bson:"scoreTypes,omitempty"` // TODO: 如果以后失败，可能是这里的原因
	RestrictedProperties   []string           `bson:"restrictedProperties,omitempty"`
	RequiredProperties     []string           `bson:"requiredProperties,omitempty"`
	PrimaryConcepts        []string           `bson:"primaryConcepts,omitempty"`
	I18Ns                  map[string]*I18N   `bson:"i18n"` // key: ISO lang 缩写
	Position               *LevelPos          `bson:"position"`
	Rewards                []*RewardBrief     `bson:"rewards"`
	Original               primitive.ObjectID `bson:"original,omitempty"`
}

// RewardBrief 是战役关卡内的rewards字段，表示当前关卡完成后的奖励。
// 和achievements集合中的rewards不同，achievements集合中的rewards表示真正会获得的奖励；
// 而这里的rewards只和下一关的箭头指示有关
type RewardBrief struct {
	Achievement string `bson:"achievement"`
	Level       string `bson:"level,omitempty"`
	Item        string `bson:"item,omitempty"`
}

// I18N 顾名思义
type I18N struct {
	Name       string `bson:"name"`
	FullName   string `bson:"fullName,omitempty"`
	Desc       string `bson:"description"`
	LoadingTip string `bson:"loadingTip,omitempty"`
}

// LevelPos 表示关卡在campaign中的位置
type LevelPos struct {
	X float64 `bson:"x"`
	Y float64 `bson:"y"`
}

// CampaignAll 旨在构建完整的Campaign结构, 并非简要结构, 未来发现新字段可以往里面一个一个加
type CampaignAll struct {
	Slug     string                 `bson:"slug"`
	Name     string                 `bson:"name"`
	FullName string                 `bson:"fullName"`
	Campaign string                 `bson:"campaign"`
	Levels   map[string]interface{} `bson:"levels"`
}
