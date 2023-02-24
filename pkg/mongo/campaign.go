package mongo

type campaign struct {
	Slug     string                 `bson:"slug"`
	Name     string                 `bson:"name"`
	FullName string                 `bson:"fullName"`
	Levels   map[string]interface{} `bson:"levels"`
}

type levelBrief struct {
	Concepts []string         `bson:"concepts"`
	Slug     string           `bson:"slug"`
	Campaign string           `bson:"campaign"`
	Name     string           `bson:"name"`
	Desc     string           `bson:"description"`
	I18Ns    map[string]*I18N `bson:"i18n"` // key: ISO lang 缩写
}

type I18N struct {
	Name       string `bson:"name"`
	Desc       string `bson:"description"`
	LoadingTip string `bson:"loadingTip"`
}
