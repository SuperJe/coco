package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"coco/pkg/mongo"
)

var collection string

func init() {
	flag.StringVar(&collection, "collection", "", "mongo collection")
}

func main() {
	flag.Parse()
	cli, err := mongo.NewClient(&mongo.ClientConfig{
		URI:     "mongodb://127.0.0.1:27017",
		DB:      "coco",
		Timeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	_, err = cli.FindOne(context.Background(), collection, bson.M{})
	if err != nil {
		panic(err)
	}
}
