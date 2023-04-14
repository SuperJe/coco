package mongo

import (
	"coco/pkg/i18n"
	"context"
	"fmt"

	"coco/pkg/mongo/entity"
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

// GetCompletedLevels 获取用户已完成的关卡
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

// GetCampaignZhHansName 英文名转中文
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
// 查level集合, 拿出campaign的英文名字, 再查campaign集合查出中文名
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
	opts := options.Find().SetProjection(bson.D{{"terrain", 1}, {"original", 1}})
	cursor, err := c.Find(ctx, collectionLevel, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "Find err")
	}
	var results []*entity.Level
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "cursor.All err")
	}

	names := make([]string, 0)
	for _, result := range results {
		names = append(names, result.CampaignName)
	}
	nameMap, err := c.GetCampaignZhHansName(ctx, names)
	if err != nil {
		return nil, errors.Wrap(err, "c.GetCampaignZhHansName err")
	}
	grouped := make(map[string][]string)
	for _, result := range results {
		name, ok := nameMap[result.CampaignName]
		if !ok {
			return nil, errors.Wrapf(err, "campaign name:%s not mapping", result.CampaignName)
		}
		grouped[name] = append(grouped[name], result.Original.Hex())
	}
	return grouped, nil
}
