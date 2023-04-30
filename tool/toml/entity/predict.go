package entity

type Config struct {
	PredictConfig *PredictConfig `toml:"predict_config"`
}

// PredictConfig predict config
type PredictConfig struct {
	// key: business
	RankParams map[string]*RankParams `toml:"rank_params"`
}

// GetRankParamItem get rank param item
func (p *PredictConfig) GetRankParamItem(business, region string) RankParamItem {
	if p == nil {
		return RankParamItem{}
	}
	if p.RankParams[business] == nil {
		return RankParamItem{}
	}
	if p.RankParams[business].Params[region] == nil {
		return RankParamItem{}
	}
	return *p.RankParams[business].Params[region]
}

// RankParams rank param
type RankParams struct {
	// key: region
	Params map[string]*RankParamItem `toml:"params"`
}

// RankParamItem rank param item
type RankParamItem struct {
	Model                   string `toml:"model"`
	Truncate                bool   `toml:"truncate"`
	FeatureAdditionalSchema string `toml:"feature_additional_schema"`
}
