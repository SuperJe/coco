package store

import (
	"context"

	"github.com/pkg/errors"
	"xorm.io/xorm"

	"github.com/SuperJe/coco/pkg/mysql"
)

// Store 存储对象
type Store struct {
	mysql *xorm.Engine
}

// NewStore 新建存储对象
func NewStore() (*Store, error) {
	cli, err := mysql.Engine()
	if err != nil {
		return nil, errors.Wrap(err, "DSEngine err")
	}
	return &Store{mysql: cli}, nil
}

func (s *Store) AddReservation(ctx context.Context, reserve *Reserve) error {
	_, err := s.mysql.Context(ctx).Insert(reserve)
	return err
}

func (s *Store) GetReservation(ctx context.Context, name string) (*Reserve, error) {
	r := &Reserve{Name: name}
	exist, err := s.mysql.Context(ctx).Get(r)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return r, nil
}
