package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/SuperJe/coco/pkg/mongo/entity"
	"github.com/SuperJe/coco/pkg/util/sliceutil"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type copyResult struct {
	srcID string
	dstID string
}

// 拷贝关卡, 返回源文档关的original和目标文档的的字段
func (s *Service) copyLevel() (*copyResult, error) {
	// 从源集合中拷贝指定关卡配置到目标集合的某个已存在的关卡上
	srcSlug := s.config.SrcSlug
	dstSlug := s.config.DstSlug
	srcDoc := bson.M{}
	if err := s.mgo.FindOne(context.Background(), levelCollectionSrc, bson.M{"slug": srcSlug}, &srcDoc); err != nil {
		return nil, errors.Wrapf(err, "Find %s err", srcSlug)
	}
	dstDoc := bson.M{}
	if err := s.mgo.FindOne(context.Background(), levelCollectionDst, bson.M{"slug": dstSlug}, &dstDoc); err != nil {
		return nil, errors.Wrapf(err, "Find %s err", dstSlug)
	}

	// 将源文档的字段复制到目标文档中, 如果目标文档已存在该字段, 则会被覆盖
	// 会保留目标文档中有的, 源文档没有的, 同时再keepFields中的字段不会被覆盖。
	for key, value := range srcDoc {
		if !sliceutil.Contain(s.config.KeepFields, key) {
			dstDoc[key] = value
			// fmt.Printf("key: %s \t\t, val:%+v\n", key, value)
		}
	}
	err := s.mgo.UpdateOne(context.Background(), levelCollectionDst, bson.M{"slug": dstSlug}, bson.M{"$set": dstDoc})
	if err != nil {
		return nil, errors.Wrap(err, "UpdateOne err")
	}
	return &copyResult{
		srcID: srcDoc["original"].(primitive.ObjectID).Hex(),
		dstID: dstDoc["original"].(primitive.ObjectID).Hex(),
	}, nil
}

func (s *Service) insertLevelToCampaign(srcID, dstID string) error {
	srcLevel, ok := s.campaignEntity.Levels[srcID]
	if !ok {
		return errors.Errorf("srcID not exist:%s", srcID)
	}
	s.campaignEntity.Levels[dstID] = srcLevel
	err := s.mgo.UpdateOne(context.Background(), "campaigns",
		bson.M{"name": s.config.Campaign}, bson.M{"$set": bson.M{"levels": s.campaignEntity.Levels}})
	if err != nil {
		return errors.Wrap(err, "UpdateOne levels err")
	}
	i18n := make(map[string]*entity.I18N)
	i18n["zh-HANS"] = &entity.I18N{Name: s.config.DstName, Desc: s.config.Desc, LoadingTip: s.config.Tip}
	if err := s.updateLevelField(dstID, "i18n", i18n); err != nil {
		return errors.Wrap(err, "s.updateLevelField err")
	}
	if err := s.updateLevelField(dstID, "slug", s.config.DstSlug); err != nil {
		return errors.Wrap(err, "s.updateLevelField err")
	}
	if err := s.updateLevelField(dstID, "position", s.config.Position); err != nil {
		return errors.Wrap(err, "s.updateLevelField err")
	}
	return nil
}

func (s *Service) updateLevelField(levelID, key string, val interface{}) error {
	field := fmt.Sprintf("levels.%s.%s", levelID, key)
	return s.mgo.UpdateOne(context.Background(), "campaigns",
		bson.M{"name": s.config.Campaign}, bson.M{"$set": bson.M{field: val}})
}

func main() {
	flag.Parse()
	s, err := newService()
	if err != nil {
		fmt.Printf("newService err:%s\n", err.Error())
		panic(err)
	}
	switch method {
	case "copy-level":
		result, err := s.copyLevel()
		if err != nil {
			fmt.Println("err:", err.Error())
			panic(err)
		}
		if err := s.insertLevelToCampaign(result.srcID, result.dstID); err != nil {
			fmt.Println("err:", err.Error())
			panic(err)
		}
	case "print-config":
		c, err := configLoad(configPath)
		fmt.Printf("config:%+v, err:%+v\n", c, err)
	default:
		fmt.Println("invalid method")
	}
}
