package model

import "encoding/json"

type Progression struct {
	Done       int32 `json:"done"`
	Total      int32 `json:"total"`
	Unfinished int32 `json:"unfinished"`
}

type CampaignProgression struct {
	Dungeon  *Progression `json:"dungeon"`
	Forest   *Progression `json:"forrest"`
	Desert   *Progression `json:"desert"`
	Mountain *Progression `json:"mountain"`
	Glacier  *Progression `json:"glacier"`
}

func (cp *CampaignProgression) FromDB(bs []byte) error {
	return json.Unmarshal(bs, cp)
}

func (cp *CampaignProgression) ToDB() ([]byte, error) {
	return json.Marshal(cp)
}
