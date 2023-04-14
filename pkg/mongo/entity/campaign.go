package entity

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
	Concepts []string         `bson:"concepts"`
	Slug     string           `bson:"slug"`
	Campaign string           `bson:"campaign"`
	Name     string           `bson:"name"`
	Desc     string           `bson:"description"`
	I18Ns    map[string]*I18N `bson:"i18n"` // key: ISO lang 缩写
}

// I18N 顾名思义
type I18N struct {
	Name       string `bson:"name"`
	FullName   string `bson:"fullName,omitempty"`
	Desc       string `bson:"description"`
	LoadingTip string `bson:"loadingTip,omitempty"`
}

// CampaignAll 旨在构建完整的Campaign结构, 并非简要结构, 未来发现新字段可以往里面一个一个加
type CampaignAll struct {
	Slug     string                 `bson:"slug"`
	Name     string                 `bson:"name"`
	FullName string                 `bson:"fullName"`
	Levels   map[string]interface{} `bson:"levels"`
}
