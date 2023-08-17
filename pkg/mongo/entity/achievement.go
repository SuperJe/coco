package entity

// EarnAchievements 是earneachievements集合的实例
type EarnAchievements struct {
	User            string         `bson:"user"`
	AchievementID   string         `bson:"achievement"`
	AchievementName string         `bson:"achievementName"`
	Points          int64          `bson:"earnedPoints"`
	Gems            int64          `bson:"earnedGems"`
	Rewards         *EarnedRewards `bson:"earnedRewards"`
}

// EarnedRewards 是获取到的奖励, users和achievements集合中都有这个结构
type EarnedRewards struct {
	Items  []string `bson:"items,omitempty" json:"items"`
	Levels []string `bson:"levels,omitempty" json:"levels"`
	Gems   int64    `bson:"gems,omitempty" json:"gems"`
}

// Achievements 是achievements集合的实例
type Achievements struct {
	Slug    string         `bson:"slug"`
	Name    string         `bson:"name"`
	Related string         `bson:"related"` // 是关卡的originalID
	Rewards *EarnedRewards `bson:"rewards"`
}
