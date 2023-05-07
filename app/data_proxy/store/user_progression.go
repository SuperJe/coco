package store

import (
	"fmt"
	"github.com/SuperJe/coco/app/data_proxy/model"
)

type UserProgression struct {
	ID        int64                      `xorm:"id"`
	Name      string                     `xorm:"name"`
	LastLevel string                     `xorm:"last_level"`
	Completed int32                      `xorm:"completed"`
	Gems      int32                      `xorm:"gems"`
	Detail    *model.CampaignProgression `xorm:"detail"`
}

func NewUserProgression(req *model.UpdateUserProgressionReq) *UserProgression {
	return &UserProgression{
		Name:      req.Name,
		LastLevel: req.LastLevel,
		Completed: req.Completed,
		Gems:      req.Gems,
		Detail:    req.CampProgression,
	}
}

type UserProgressions []*UserProgression

func (ups UserProgressions) GroupByName() map[string]*model.CampaignProgression {
	result := make(map[string]*model.CampaignProgression, len(ups))
	for _, up := range ups {
		up := up
		result[up.Name] = up.Detail
	}
	return result
}

func (up UserProgression) TableName() string {
	return "user_progression"
}

// UpsertUserProgression 存在则更新 不存在则插入
func (s *Store) UpsertUserProgression(up *UserProgression) error {
	if up == nil {
		return fmt.Errorf("UserProgression nil")
	}
	cond := &UserProgression{Name: up.Name}
	exist, err := s.mysql.Exist(cond)
	if err != nil {
		return err
	}
	if !exist {
		if _, err := s.mysql.Insert(up); err != nil {
			return err
		}
		return nil
	}
	if _, err := s.mysql.Update(up, cond); err != nil {
		return err
	}
	return nil

}

func (s *Store) GetUserProgression(name string) (*UserProgression, error) {
	up := &UserProgression{Name: name}
	exist, err := s.mysql.Get(up)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return up, nil
}

func (s *Store) BatchGetUserProgressions(names []string) (UserProgressions, error) {
	progressions := make([]*UserProgression, 0)
	if err := s.mysql.Table(UserProgression{}.TableName()).In("name", names).Find(&progressions); err != nil {
		return nil, err
	}
	fmt.Printf("len:%d\n", len(progressions))
	for _, progression := range progressions {
		fmt.Printf("p:%+v\n", progression)
	}
	return progressions, nil
}
