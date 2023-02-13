package mongo

import (
	"context"
	"time"

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
