package service

import (
	"github.com/SuperJe/coco/app/data_proxy/store"
	"github.com/pkg/errors"
)

// Service 服务
type Service struct {
	store *store.Store
}

// NewService 新建service
func NewService() (*Service, error) {
	s, err := store.NewStore()
	if err != nil {
		return nil, errors.Wrap(err, "store.NewStore err")
	}
	return &Service{store: s}, nil
}
