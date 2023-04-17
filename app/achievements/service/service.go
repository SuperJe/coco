package service

import (
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/mongo/entity"
)

type Service struct {
	mgo *mongo.Client
	earned *entity.EarnAchievements
}
