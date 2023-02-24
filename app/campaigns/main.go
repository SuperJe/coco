package main

import (
	"context"
	"fmt"

	"coco/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var mgo *mongo.Client

func main() {
	var err error
	mgo, err = mongo.NewCocoClient()
	if err != nil {
		panic(err)
	}

	str, err := mgo.FindOne(context.Background(), "campaigns", bson.M{"name": "Dungeon"})
	if err != nil {
		fmt.Println("FindOne err:", err.Error())
		return
	}
	fmt.Printf("find one:%s\n", str.(string))
}
