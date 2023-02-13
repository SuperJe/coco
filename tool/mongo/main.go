package main

import (
	"context"
	"time"

	"coco/pkg/mongo"
)

func main() {
	cli, err := mongo.NewClient(&mongo.ClientConfig{
		URI:     "mongodb://localhost:27017",
		DB:      "coco",
		Timeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	_, err = cli.Find(context.Background())
	if err != nil {
		panic(err)
	}
}
