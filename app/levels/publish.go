package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SuperJe/coco/pkg/i18n"
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/mongo/entity"
	"github.com/SuperJe/coco/pkg/util"
)

var (
	// required
	camp             string
	pubOriginal      string
	preOriginal      string
	nextOriginal     string
	preAchievementID string
	pubAchievementID string
	// optional
	posX, posY                                           float64
	levelName, levelFullName, levelDesc, levelLoadingTip string
)

func init() {
	flag.StringVar(&camp, "camp", "", "camp name")
	flag.StringVar(&pubOriginal, "pub_original", "", "publish level original id")
	flag.StringVar(&preOriginal, "pre_original", "", "pre level original id")
	flag.StringVar(&nextOriginal, "next_original", "", "next level original id")
	flag.StringVar(&preAchievementID, "pre_acv", "", "pre achievement id")
	flag.StringVar(&pubAchievementID, "pub_acv", "", "publish level achievement id")
	flag.StringVar(&levelName, "name", "颜总监真牛", "level name")
	flag.StringVar(&levelFullName, "full_name", "颜总监真棒", "level full name")
	flag.StringVar(&levelDesc, "desc", "颜总监真帅", "level desc")
	flag.StringVar(&levelLoadingTip, "loading_tip", "颜总监泰裤辣", "level loading tip")
	flag.Float64Var(&posX, "pos_x", 37.0, "position x, float64")
	flag.Float64Var(&posY, "pos_y", 37.0, "position y, float64")
}

func paramCheck() {
	if util.EmptyS(camp) || util.EmptyS(pubOriginal) || util.EmptyS(preOriginal) ||
		util.EmptyS(nextOriginal) || util.EmptyS(preAchievementID) || util.EmptyS(pubAchievementID) {
		panic("empty field")
	}
}

func (s *Service) publish(ctx context.Context) error {
	paramCheck()
	// 从level集合中找到对应文档
	level := &entity.Level{}
	objID, err := primitive.ObjectIDFromHex(pubOriginal)
	if err != nil {
		return errors.Wrapf(err, "%s ObjectIDFromHex err", pubOriginal)
	}
	if err = s.mgo.FindOne(ctx, mongo.CollectionLevel, bson.M{"original": objID}, level); err != nil {
		return errors.Wrap(err, "FindOne err")
	}
	// 自定义关卡名字信息、所属战役以及奖励信息
	levelBrief := level.ToBrief()
	levelBrief.Position = &entity.LevelPos{X: posX, Y: posY}
	levelBrief.I18Ns = map[string]*entity.I18N{i18n.ZhHans: &entity.I18N{
		Name:       levelName,
		FullName:   levelFullName,
		Desc:       levelDesc,
		LoadingTip: levelLoadingTip,
	}}
	levelBrief.Rewards = append(levelBrief.Rewards, &entity.RewardBrief{
		Achievement: pubAchievementID,
		Level:       nextOriginal,
	})
	levelBrief.Campaign = strings.ToLower(camp)
	fmt.Printf("level brief:\n%s\n", util.JSONString(levelBrief))
	// 插入campaign中
	if err = s.mgo.PublishLevel(ctx, levelBrief); err != nil {
		return errors.Wrap(err, "UpdateOne levels err")
	}
	fmt.Println("insert success")
	// 更新奖励信息:
	// 1. 将上一关卡的战役rewards指向当前关卡, 当前关卡的战役rewards指向下一关卡
	// 2. 将上一关卡的achievements中的rewards的level字段更新成当前关卡, 同理, 当前关卡相应字段更新为下一关卡
	// 但是当前关卡的战役rewards在上面publish的时候已经更新;
	// 当前关卡的achievements中的rewards的level字段在关卡编辑器的成就中已经编辑，无需再更新
	// 所以只需要更新上一关卡指向当前关卡即可
	return s.mgo.LinkRewards(ctx, camp, preOriginal, pubOriginal, preAchievementID)
}
