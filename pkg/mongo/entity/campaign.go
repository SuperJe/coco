package entity

// Campaign 战役, 如Dungeon, Desert等大战役
type Campaign struct {
	Slug     string                 `bson:"slug"`
	Name     string                 `bson:"name"`
	FullName string                 `bson:"fullName"`
	Levels   map[string]*LevelBrief `bson:"levels"`
}

// LevelBrief 在战役中有非常多的小关卡, 此为保存在战役结构中的小关卡的关键数据结构
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
	Desc       string `bson:"description"`
	LoadingTip string `bson:"loadingTip"`
}
