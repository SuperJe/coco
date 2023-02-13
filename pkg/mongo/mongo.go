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

// NewClient new mongo client
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

func (c *Client) Find(ctx context.Context) (interface{}, error) {
	cursor, err := c.cli.Database(c.db).Collection("users").Find(ctx, bson.M{"name": "codeMagic"})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Println("cursor close err:", err.Error())
		}
	}()

	fmt.Println("current cursor:", cursor.Current.String())
	return nil, nil
}
