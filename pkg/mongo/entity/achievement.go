package entity

type EarnAchievements struct {
	User            string         `bson:"user"`
	AchievementID   string         `bson:"achievement"`
	AchievementName string         `bson:"achievementName"`
	Points          int64          `bson:"earnedPoints"`
	Gems            int64          `bson:"earnedGems"`
	Rewards         *EarnedRewards `bson:"earnedRewards"`
}

type EarnedRewards struct {
	Items  []string `bson:"items"`
	Levels []string `bson:"levels"`
	Gems   int64    `bson:"gems"`
}
