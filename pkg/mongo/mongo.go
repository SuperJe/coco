package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientConfig struct {
	URI     string
	DB      string // default DB
	Timeout time.Duration
}

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

func (c *Client) FindOne(ctx context.Context, collection string, filter bson.M) (interface{}, error) {
	r := c.cli.Database(c.db).Collection(collection).FindOne(ctx, filter)
	if r.Err() != nil {
		fmt.Printf("find err: %s, db:%s\n", r.Err().Error(), c.db)
		return nil, r.Err()
	}
	raw, err := r.DecodeBytes()
	if err != nil {
		fmt.Printf("r.DecodeBytes err:%s\n", err.Error())
		return nil, err
	}
	fmt.Printf("raw string:%s\n", raw.String())
	return raw.String(), nil
}
