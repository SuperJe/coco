package main

import (
	"coco/pkg/mongo/entity"
	"context"
	"encoding/json"
	"fmt"

	"coco/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var mgo *mongo.Client

const (
	collection = "campaigns"
)

func main() {
	var err error
	mgo, err = mongo.NewCocoClient()
	if err != nil {
		panic(err)
	}

	dungeon := &entity.Campaign{}
	if err := mgo.FindOne(context.Background(), collection, bson.M{"name": "Dungeon"}, dungeon); err != nil {
		fmt.Println("FindOne err:", err.Error())
		return
	}
	bs, _ := json.Marshal(dungeon)
	fmt.Printf("dungeon:\n\n:%s", string(bs))
}
