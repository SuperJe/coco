package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Progression struct {
	Done       int32 `json:"done"`
	Total      int32 `json:"total"`
	Unfinished int32 `json:"unfinished"`
}

func (p Progression) String() string {
	return fmt.Sprintf("进度:%d/%d", p.Done, p.Total)
}

type CampaignProgression struct {
	Dungeon  *Progression `json:"dungeon"`
	Forest   *Progression `json:"forrest"`
	Desert   *Progression `json:"desert"`
	Mountain *Progression `json:"mountain"`
	Glacier  *Progression `json:"glacier"`
}

func (cp *CampaignProgression) String() string {
	values := make([]string, 0, 5)
	values = append(values, fmt.Sprintf("地牢%s", cp.Dungeon.String()))
	values = append(values, fmt.Sprintf("森林%s", cp.Forest.String()))
	values = append(values, fmt.Sprintf("沙漠%s", cp.Desert.String()))
	values = append(values, fmt.Sprintf("山峰%s", cp.Mountain.String()))
	values = append(values, fmt.Sprintf("冰川%s", cp.Mountain.String()))
	return strings.Join(values, ",")
}

func (cp *CampaignProgression) FromDB(bs []byte) error {
	return json.Unmarshal(bs, cp)
}

func (cp *CampaignProgression) ToDB() ([]byte, error) {
	return json.Marshal(cp)
}
