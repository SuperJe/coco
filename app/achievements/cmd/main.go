package cmd

import (
	"coco/pkg/mongo"
	"coco/pkg/mongo/entity"
)

var (
	mgo    *mongo.Client
	earned = &entity.EarnAchievements{}
)

func init() {
	var err error
	mgo, err = mongo.NewCocoClient()
	if err != nil {
		panic(err)
	}
}

func main() {

}
