package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClientConfig 配置文件
type ClientConfig struct {
	URI     string
	DB      string // default DB
	Timeout time.Duration
}

// Client 客户端
type Client struct {
	db  string
	cli *mongo.Client
}

// NewClient 通用mongo client
func NewClient(c *ClientConfig) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI))
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return &Client{db: c.DB, cli: cli}, err
}

// NewCocoClient 本机coco数据库的mongo client
func NewCocoClient() (*Client, error) {
	c := &ClientConfig{
		URI:     "mongodb://127.0.0.1:27017",
		DB:      "coco",
		Timeout: time.Second * 5,
	}
	return NewClient(c)
}

// NewCocoClient2 宿主机连接coco数据库的mongo client
// 在容器中暴露27017的端口映射到27018的宿主机上, 所以这里访问27018端口
// 来实现在宿主机上访问容器内的mongo
func NewCocoClient2() (*Client, error) {
	c := &ClientConfig{
		URI:     "mongodb://127.0.0.1:27018",
		DB:      "coco",
		Timeout: time.Second * 5,
	}
	return NewClient(c)
}

// Find 查找全部
func (c *Client) Find(ctx context.Context, collection string,
	filter bson.M, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return c.Collection(collection).Find(ctx, filter, opts...)
}

// FindOne 匹配一条数据
func (c *Client) FindOne(ctx context.Context, collection string,
	filter bson.M, v interface{}, opts ...*options.FindOneOptions) error {
	r := c.Collection(collection).FindOne(ctx, filter, opts...)
	if r.Err() != nil {
		fmt.Printf("find err: %s, db:%s\n", r.Err().Error(), c.db)
		return r.Err()
	}
	return r.Decode(v)
}

// UpdateOne 更新一条数据
func (c *Client) UpdateOne(ctx context.Context, collection string, filter, updater interface{}) error {
	result, err := c.cli.Database(c.db).Collection(collection).UpdateOne(ctx, filter, updater)
	if err != nil {
		return err
	}
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
	}
	return nil
}

// InsertOne 插入一条数据
func (c *Client) InsertOne(ctx context.Context, collection string, doc interface{}) error {
	res, err := c.Collection(collection).InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	fmt.Println("insertID:", res.InsertedID)
	return nil
}

// Collection 指定集合
func (c *Client) Collection(collection string) *mongo.Collection {
	return c.cli.Database(c.db).Collection(collection)
}
