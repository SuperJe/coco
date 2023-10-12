package service

import "github.com/SuperJe/coco/app/home/internal/store"

type Service struct {
	store *store.Store
}

func NewService() (*Service, error) {
	s, err := store.NewStore()
	if err != nil {
		return nil, err
	}
	return &Service{store: s}, nil
}
