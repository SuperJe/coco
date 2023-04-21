package store

import (
	"github.com/SuperJe/coco/pkg/mysql"
	"github.com/pkg/errors"
)

// Store 存储对象
type Store struct {
	mysql *mysql.Client
}

// NewStore 新建存储对象
func NewStore() (*Store, error) {
	cli, err := mysql.DSEngine()
	if err != nil {
		return nil, errors.Wrap(err, "DSEngine err")
	}
	return &Store{mysql: cli}, nil
}
