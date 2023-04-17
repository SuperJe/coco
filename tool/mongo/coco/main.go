package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/util"
)

var cli *mongo.Client

const codeMagic = "codeMagic"

type Progression struct {
	Done       int32 `json:"done"`
	Unfinished int32 `json:"unfinished"`
	Total      int32 `json:"total"`
}

type CampaignProgression struct {
	Dungeon  *Progression `json:"dungeon"`
	Forest   *Progression `json:"forrest"`
	Desert   *Progression `json:"desert"`
	Mountain *Progression `json:"mountain"`
	Glacier  *Progression `json:"glacier"`
}

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

func countsLevel() {
	counts, err := cli.CountLevels(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("counts:", util.JSONString(counts))
}

func buildProgression(campaign string, completed map[string][]string, counts map[string]int32) *Progression {
	return &Progression{
		Done:       int32(len(completed[campaign])),
		Unfinished: counts[campaign] - int32(len(completed[campaign])),
		Total:      counts[campaign],
	}
}

func getCampProgressions() {
	start := time.Now()
	counts, err := cli.CountLevels(context.Background())
	if err != nil {
		panic(err)
	}
	levels, err := cli.GetCompletedLevels(context.Background(), codeMagic)
	if err != nil {
		panic(err)
	}
	completed, err := cli.GroupLevelByCampaign(context.Background(), levels)
	if err != nil {
		panic(err)
	}
	progressions := &CampaignProgression{}
	progressions.Dungeon = buildProgression("Dungeon", completed, counts)
	progressions.Forest = buildProgression("Forest", completed, counts)
	progressions.Desert = buildProgression("Desert", completed, counts)
	progressions.Mountain = buildProgression("Mountain", completed, counts)
	progressions.Glacier = buildProgression("Glacier", completed, counts)
	fmt.Printf("CampaignProgression:%s\n", util.JSONString(progressions))
	fmt.Printf("cost:%+v\n", time.Since(start))
}

func main() {
	var err error
	cli, err = mongo.NewCocoClient2()
	if err != nil {
		panic(err)
	}
	// getUserByName()
	// getEarnedLevels()
	// nameMapping()
	// countsLevel()
	getCampProgressions()
}
