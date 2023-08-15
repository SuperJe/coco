package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) updateLevelField(levelID, key string, val interface{}) error {
	field := fmt.Sprintf("levels.%s.%s", levelID, key)
	return s.mgo.UpdateOne(context.Background(), "campaigns",
		bson.M{"name": s.config.Campaign}, bson.M{"$set": bson.M{field: val}})
}
