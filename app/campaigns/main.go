package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"coco/pkg/mongo"
	"coco/pkg/mongo/entity"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	method             string
	selectedLevelsFile string
	mgo                *mongo.Client
	dungeon            = &entity.Campaign{}
)

const (
	collection            = "campaigns"
	levelWithIDFile       = "/home/coco/codecombat/data/coco/doc/campaign/level_with_id.txt"
	methodSelectedLevelID = "selected_level_id"
)

// 建立level ObjectId -> level中文名的映射
func getLevelMapping() map[string]string {
	mapping := make(map[string]string, len(dungeon.Levels))
	for id, level := range dungeon.Levels {
		i18n, ok := level.I18Ns["zh-HANS"]
		if !ok {
			fmt.Printf("level %s miss zh-HANS i18n\n", level.Name)
			continue
		}
		mapping[i18n.Name] = id
	}
	fmt.Println("dungeon 所有关卡数:", len(mapping))
	return mapping
}

func getSelectedLevels() ([]string, error) {
	// 文件不大, 可以一次性读取进内存
	bs, err := ioutil.ReadFile(selectedLevelsFile)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bs), "\n"), nil
}

func writeSelectedLevelsWithID() error {
	if len(selectedLevelsFile) == 0 {
		return fmt.Errorf("请指定文件名\n")
	}
	mapping := getLevelMapping()
	names, err := getSelectedLevels()
	if err != nil {
		return err
	}

	// 找到需要的关卡id, 一起写入新文件
	total := 0
	buff := &bytes.Buffer{}
	for _, name := range names {
		id, ok := mapping[name]
		if !ok {
			continue
		}
		str := fmt.Sprintf("id\t%s\tname\t%s\n", id, name)
		if _, err := buff.WriteString(str); err != nil {
			return err
		}
		total++
	}
	if err := ioutil.WriteFile(levelWithIDFile, buff.Bytes(), fs.FileMode(0666)); err != nil {
		return err
	}
	fmt.Println("写入路径:", levelWithIDFile)
	fmt.Println("成功写入带id的关卡数量:", total)
	return nil
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
	flag.StringVar(&method, "method", "", "执行方法:  selected_level_id")
	flag.StringVar(&selectedLevelsFile, "level_file", "", "需要关卡的文件名")
}

func main() {
	flag.Parse()
	switch method {
	case methodSelectedLevelID:
		if err := writeSelectedLevelsWithID(); err != nil {
			panic(err)
		}
	default:
		panic("method invalid")
	}
}
