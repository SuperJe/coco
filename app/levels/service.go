package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/mongo/entity"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	method     string
	configPath string
)

const (
	levelCollectionSrc = "levelstmp"
	levelCollectionDst = "levels"
)

func init() {
	// keepFields = []string{"_id", "original", "index", "slug", "creator", "watchers", "permissions", "name"}
	flag.StringVar(&method, "method", "", "method: copy-level")
	flag.StringVar(&configPath, "config", "", "config file path")
}

type Service struct {
	config         *Config
	mgo            *mongo.Client
	campaignEntity *entity.CampaignAll
}

func newService() (*Service, error) {
	c, err := configLoad(configPath)
	if err != nil {
		return nil, err
	}
	if len(c.SrcSlug) == 0 || len(c.DstSlug) == 0 || len(c.Campaign) == 0 {
		return nil, fmt.Errorf("config err:\n%+v", c)
	}
	mgo, err := mongo.NewCocoClient2()
	if err != nil {
		panic(err)
	}
	campaign := &entity.CampaignAll{}
	if err := mgo.FindOne(context.Background(), "campaigns", bson.M{"name": c.Campaign}, campaign); err != nil {
		fmt.Printf("FindOne err:%s, campaign:%s\n", err.Error(), campaign)
		panic(err)
	}
	return &Service{
		config:         c,
		mgo:            mgo,
		campaignEntity: campaign,
	}, nil

}

type Config struct {
	SrcSlug    string    `toml:"src_slug"`
	DstSlug    string    `toml:"dst_slug"`
	Campaign   string    `toml:"campaign_name"`
	KeepFields []string  `toml:"keep_fields"`
	DstName    string    `toml:"dst_name"`
	Desc       string    `toml:"description"`
	Tip        string    `toml:"loading_tip"`
	Position   *Position `toml:"position" bson:"position"`
}

type Position struct {
	X int `toml:"x" bson:"x"`
	Y int `toml:"y" bson:"y"`
}

func configLoad(file string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(file, c); err != nil {
		return nil, err
	}
	return c, nil
}
