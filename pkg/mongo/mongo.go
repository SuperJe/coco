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

type User struct {
	CreateHost string `bson:"createdOnHost"`
	Anonymous  bool   `bson:"anonymous"`
	Email      string `bson:"email"`
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

func (c *Client) Find(ctx context.Context, collection string) (interface{}, error) {
	filter := bson.M{"name": "codeMagic"}
	r := c.cli.Database(c.db).Collection(collection).FindOne(ctx, filter)
	if r.Err() != nil {
		fmt.Printf("find err: %s, db:%s\n", r.Err().Error(), c.db)
		return nil, r.Err()
	}
	// raw, err := r.DecodeBytes()
	user := &User{}
	if err := r.Decode(user); err != nil {
		fmt.Printf("r.Decode err:%s", err.Error())
		return nil, err
	}
	fmt.Printf("user:%+v", user)
	// if err != nil {
	// 	fmt.Printf("r.DecodeBytes err:%s\n", err.Error())
	// 	return nil, err
	// }
	// fmt.Printf("raw string:%s\n", raw.String())
	return nil, nil
}
