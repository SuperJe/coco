package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

type Config struct {
	User     string
	Password string
	DB       string
}

type Client struct {
	engine *xorm.Engine
}

// DSEngine data_sync数据库的engine
func DSEngine() (*Client, error) {
	c := &Config{
		User:     "root",
		Password: "123456",
		DB:       "data_sync",
	}
	return NewEngine(c)
}

// NewEngine 根据配置新建mysql引擎
func NewEngine(c *Config) (*Client, error) {
	dsn := fmt.Sprintf("%s:%s/%s?charset=utf8", c.User, c.Password, c.DB)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "xorm.NewEngine err")
	}
	if err := engine.Ping(); err != nil {
		return nil, errors.Wrap(err, "Ping err")
	}
	return &Client{engine: engine}, nil
}
