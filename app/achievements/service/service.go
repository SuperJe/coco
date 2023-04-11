package service

import (
	"coco/pkg/mongo"
	"coco/pkg/mongo/entity"
)

type Service struct {
	mgo    *mongo.Client
	earned *entity.EarnAchievements
}
