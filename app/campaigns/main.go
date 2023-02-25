package main

import (
	"context"
	"flag"
	"fmt"

	"coco/pkg/mongo"
	"coco/pkg/mongo/entity"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	method  string
	mgo     *mongo.Client
	dungeon = &entity.Campaign{}
)

const (
	collection        = "campaigns"
	methodMappingName = "mapping_name"
)

// 建立level ObjectId -> level中文名的映射
func mapLevelName() {
	mapping := make(map[string]string, len(dungeon.Levels))
	for id, level := range dungeon.Levels {
		i18n, ok := level.I18Ns["zh-HANS"]
		if !ok {
			fmt.Printf("level %s miss zh-HANS i18n\n", level.Name)
			continue
		}
		mapping[id] = i18n.Name
	}
	fmt.Println("levels num:", len(mapping))
	for id, name := range mapping {
		fmt.Printf("id: %s,    name:%s\n", id, name)
	}
}

// 进程初始化
func init() {
	var err error
	mgo, err = mongo.NewCocoClient()
	if err != nil {
		panic(err)
	}
	if err := mgo.FindOne(context.Background(), collection, bson.M{"name": "Dungeon"}, dungeon); err != nil {
		fmt.Println("FindOne err:", err.Error())
		panic(err)
	}
	flag.StringVar(&method, "method", "", "执行方法")
}

func main() {
	flag.Parse()
	switch method {
	case methodMappingName:
		mapLevelName()
	default:
		panic("method invalid")
	}
}
