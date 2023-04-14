package main

import (
	"context"
	"fmt"

	"coco/pkg/mongo"
)

var cli *mongo.Client

const codeMagic = "codeMagic"

func getUserByName() {
	user, err := cli.GetUserByName(context.Background(), codeMagic)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user:%+v\n", user)
}

func getEarnedLevels() {
	levels, err := cli.GetCompletedLevels(context.Background(), codeMagic)
	if err != nil {
		panic(err)
	}
	grouped, err := cli.GroupLevelByCampaign(context.Background(), levels)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete levels:", grouped)
}

func nameMapping() {
	mapping, err := cli.GetCampaignZhHansName(context.Background(), []string{"Dungeon"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("mapping:%+v\n", mapping)
}

func main() {
	var err error
	cli, err = mongo.NewCocoClient2()
	if err != nil {
		panic(err)
	}
	// getUserByName()
	getEarnedLevels()
	// nameMapping()
}
