package mongo

import (
	"context"
	"fmt"
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
	if err := c.FindOne(ctx, collectionUser, filter, user); err != nil {
		return nil, errors.Wrap(err, "FindOne err")
	}
	return user, nil
}

// GetCompletedLevels 获取用户已完成的关卡, name是用户的user name
func (c *Client) GetCompletedLevels(ctx context.Context, name string) ([]string, error) {
	// 先查用户的id
	user, err := c.GetUserByName(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "c.GetUserByName err")
	}
	filter := bson.M{"user": user.ID.Hex()}
	cursor, err := c.Collection(collectionEarned).Find(ctx, filter)
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
	cursor, err := c.Find(ctx, collectionCamp, filter, opts)
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
	cursor, err := c.Find(ctx, collectionLevel, filter, opts)
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
	cursor, err := c.Find(ctx, collectionCamp, filter, opts)
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
