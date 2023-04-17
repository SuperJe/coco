package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/mongo/entity"
	"github.com/SuperJe/coco/pkg/util/sliceutil"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	method             string
	selectedLevelsFile string
	mgo                *mongo.Client
	dungeon            = &entity.Campaign{}
)

const (
	collection       = "campaigns"
	deletedLevelFile = "/home/coco/codecombat/data/coco/doc/campaign/deleted_level.txt"
	achievementsFile = "/home/coco/codecombat/data/coco/doc/campaign/achievements.txt"
	methodRebuildAll = "rebuild_all"
)

// 返回 level中文名 -> level ObjectId的映射
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

func writeDeletedLevelsWithID() error {
	if len(selectedLevelsFile) == 0 {
		return fmt.Errorf("请指定文件名\n")
	}
	// 找到不需要的关卡id, 写入新文件
	allLevels := getLevelMapping()
	names, err := getSelectedLevels()
	if err != nil {
		return err
	}
	selectedLevels := sliceutil.ToStringSet(names)
	total := 0
	buff := &bytes.Buffer{}
	for name, id := range allLevels {
		if _, ok := selectedLevels[name]; ok {
			continue
		}
		str := fmt.Sprintf("id\t%s\tname\t%s\n", id, name)
		if _, err := buff.WriteString(str); err != nil {
			return err
		}
		total++
	}
	if err := ioutil.WriteFile(deletedLevelFile, buff.Bytes(), fs.FileMode(0666)); err != nil {
		return err
	}
	fmt.Println("写入路径:", deletedLevelFile)
	fmt.Println("成功写入带id的关卡数量:", total)
	return nil
}

func reBuildAchievements() error {
	// 建立achievements集合的slug->重建的奖励level映射
	name2ID := getLevelMapping()
	selected, err := getSelectedLevels()
	if err != nil {
		return err
	}
	buff := &bytes.Buffer{}
	achievements := make(map[string]string, len(selected))
	for i := 0; i < len(selected); i++ {
		id := name2ID[selected[i]]
		slug := dungeon.Levels[id].Slug
		if i+1 < len(selected) {
			// 建立映射, 当前slug完成的奖励为下一关卡的id
			key := slug + "-complete"
			level := name2ID[selected[i+1]]
			achievements[key] = level
			str := fmt.Sprintf("%s:%s\n", key, level)
			if _, err := buff.WriteString(str); err != nil {
				return err
			}
			// 有可能是completed
			str = fmt.Sprintf("%sd:%s\n", key, level)
			if _, err := buff.WriteString(str); err != nil {
				return err
			}
		}
	}
	if err := ioutil.WriteFile(achievementsFile, buff.Bytes(), fs.FileMode(0666)); err != nil {
		return err
	}
	fmt.Println("重新奖励索引完成, 共重建:", len(achievements))
	return nil
}

// 重建地牢关卡
func reBuildDungeon() error {
	if err := writeDeletedLevelsWithID(); err != nil {
		return err
	}
	if err := reBuildAchievements(); err != nil {
		return err
	}
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
	case methodRebuildAll:
		if err := reBuildDungeon(); err != nil {
			panic(err)
		}
	default:
		panic("method invalid")
	}
}
