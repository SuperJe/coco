package mongo

import (
	"context"
	"fmt"
	"github.com/SuperJe/coco/pkg/util"
	"strings"

	"github.com/SuperJe/coco/pkg/i18n"
	"github.com/SuperJe/coco/pkg/mongo/entity"
	"github.com/SuperJe/coco/pkg/util/stringutil"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUserByName 获取用户资料
func (c *Client) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	filter := bson.M{"name": name}
	user := &entity.User{}
	if err := c.FindOne(ctx, CollectionUser, filter, user); err != nil {
		return nil, errors.Wrap(err, "FindOne err")
	}
	return user, nil
}

// GetUsers 获取所有用户信息, 当前只会返回name,earned和lastLevel
func (c *Client) GetUsers(ctx context.Context) ([]*entity.User, error) {
	opts := options.Find().SetProjection(bson.D{{"name", 1}, {"earned", 1}, {"lastLevel", 1}})
	cursor, err := c.Find(ctx, CollectionUser, nil, opts)
	if err != nil {
		return nil, errors.Wrap(err, "c.Find err")
	}
	var result []*entity.User
	if err := cursor.All(ctx, &result); err != nil {
		return nil, errors.Wrap(err, "cursor.All err")
	}
	return result, nil
}

// GetCompletedLevels 获取用户已完成的关卡, name是用户的user name
func (c *Client) GetCompletedLevels(ctx context.Context, name string) ([]string, error) {
	// 先查用户的id
	user, err := c.GetUserByName(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "c.GetUserByName err")
	}
	filter := bson.M{"user": user.ID.Hex()}
	cursor, err := c.Collection(CollectionEarned).Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "Find err")
	}

	var results []*entity.EarnAchievements
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "cursor.All err")
	}
	levels := make([]string, 0, len(results))
	for _, result := range results {
		if result.Rewards == nil {
			continue
		}
		levels = append(levels, result.Rewards.Levels...)
	}
	return levels, nil
}

// GetCampaignZhHansName 英文名转中文, 如Dungeon->地牢
func (c *Client) GetCampaignZhHansName(ctx context.Context, names []string) (map[string]string, error) {
	filter := bson.M{"name": bson.M{"$in": names}}
	opts := options.Find().SetProjection(bson.D{{"i18n", 1}, {"name", 1}})
	cursor, err := c.Find(ctx, CollectionCamp, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "FindOne err")
	}
	var results []*entity.Campaign
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "cursor.All err")
	}

	mapping := make(map[string]string)
	for _, result := range results {
		fmt.Printf("result:%+v\n", result)
		mapping[result.Name] = result.I18Ns[i18n.ZhHans].Name
	}
	return mapping, nil
}

// GroupLevelByCampaign 按campaign分类level
// 查level集合, 拿出campaign的英文名字, 根据英文名字来分类, 如Dungeon -> {"level1", "level2"}
func (c *Client) GroupLevelByCampaign(ctx context.Context, levels []string) (map[string][]string, error) {
	objLevels := make([]primitive.ObjectID, 0, len(levels))
	for _, level := range levels {
		objID, err := primitive.ObjectIDFromHex(level)
		if err != nil {
			return nil, errors.Wrap(err, "ObjectIDFromHex err")
		}
		objLevels = append(objLevels, objID)
	}
	filter := bson.M{"original": bson.M{"$in": objLevels}}
	opts := options.Find().SetProjection(bson.D{{"campaign", 1}, {"original", 1}})
	cursor, err := c.Find(ctx, CollectionLevel, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "Find err")
	}
	var results []*entity.Level
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "cursor.All err")
	}

	grouped := make(map[string][]string)
	for _, result := range results {
		camp := stringutil.CapFirstLetter(result.CampaignName)
		grouped[camp] = append(grouped[camp], result.Original.Hex())
	}
	return grouped, nil
}

// CountLevels 计算每个campaign的level数量
// 返回 campaign.Name -> 数量, 如Dungeon->56
func (c *Client) CountLevels(ctx context.Context) (map[string]int32, error) {
	camps := []string{"Dungeon", "Forest", "Desert", "Mountain", "Glacier"}
	filter := bson.M{"name": bson.M{"$in": camps}}
	opts := options.Find().SetProjection(bson.D{{"name", 1}, {"levels", 1}})
	cursor, err := c.Find(ctx, CollectionCamp, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "c.Find err")
	}
	var result []*entity.Campaign
	if err := cursor.All(ctx, &result); err != nil {
		return nil, errors.Wrap(err, "cursor.All err")
	}

	counts := make(map[string]int32, len(camps))
	for _, r := range result {
		counts[r.Name] = int32(len(r.Levels))
	}
	return counts, nil
}

// RegisterUser 新建一个用户
// 用户名前导和尾部不能是空格；中间不能含有连续空格；所有字符只能是数字+字母+空格的组合。
func (c *Client) RegisterUser(ctx context.Context, name, pwd string) error {
	if util.EmptyS(name) || util.EmptyS(pwd) {
		return fmt.Errorf("invalid params, name:%s, pwd:%s", name, pwd)
	}
	if name[0] == ' ' || name[len(name)-1] == ' ' {
		return fmt.Errorf("user name can not begin or end with space")
	}
	// 用户名只能是空格, 数字和字母的组合, 且不能有连续的空格
	spacePos := -2
	for i, ch := range name {
		if ch == ' ' {
			if spacePos+1 == i {
				return fmt.Errorf("user name can not contain consecutive spaces")
			}
			spacePos = i
			continue
		}
		if ch >= '0' && ch <= '9' {
			continue
		}
		if ch >= 'a' && ch <= 'z' {
			continue
		}
		if ch >= 'A' && ch <= 'Z' {
			continue
		}
		return fmt.Errorf("unsupported character:%c", ch)
	}
	user := entity.NewUser(name, pwd)
	if err := c.InsertOne(ctx, CollectionUser, user); err != nil {
		return errors.Wrap(err, "c.InsertOne err")
	}
	return nil
}

// PublishLevel 将已经编辑好的关卡发布到campaign中
func (c *Client) PublishLevel(ctx context.Context, level *entity.LevelBrief) error {
	field := "levels." + level.Original.Hex()
	camp := stringutil.CapFirstLetter(level.Campaign)
	err := c.UpdateOne(ctx, CollectionCamp, bson.M{"name": camp}, bson.M{"$set": bson.M{field: level}})
	if err != nil {
		return errors.Wrap(err, "c.UpdateOne err")
	}
	// 更新关卡所属的的campaign
	return c.UpdateOne(ctx, CollectionLevel, bson.M{"name": level.Name},
		bson.M{"$set": bson.M{"campaign": strings.ToLower(camp)}})

}

// LinkRewards 将pre关卡的奖励更新为next
func (c *Client) LinkRewards(ctx context.Context, camp, preLevel, curLevel, preAcv string) error {
	if len(preLevel) == 0 || len(curLevel) == 0 || len(preAcv) == 0 || len(camp) == 0 {
		return fmt.Errorf("exist empty field")
	}
	objID, err := primitive.ObjectIDFromHex(preAcv)
	if err != nil {
		return errors.Wrap(err, "primitive.ObjectIDFromHex err")
	}
	// 更新战役中的rewards: 上一个关卡的奖励指向当前
	field := fmt.Sprintf("levels.%s.rewards", preLevel)
	rbs := make([]*entity.RewardBrief, 0)
	rbs = append(rbs, &entity.RewardBrief{
		Achievement: preAcv,
		Level:       curLevel,
	})
	cp := stringutil.CapFirstLetter(camp)
	err = c.UpdateOne(ctx, CollectionCamp, bson.M{"name": cp}, bson.M{"$set": bson.M{field: rbs}})
	if err != nil {
		return errors.Wrap(err, "c.UpdateOne err")
	}
	// 更新achievements中的rewards: 上一个关卡奖励指向当前
	field = "rewards.levels"
	val := []string{curLevel}
	err = c.UpdateOne(ctx, CollectionAcv, bson.M{"_id": objID}, bson.M{"$set": bson.M{field: val}})
	if err != nil {
		return errors.Wrap(err, "c.UpdateOne err")
	}
	return nil
}
