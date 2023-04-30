package main

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/SuperJe/coco/tool/toml/entity"
)

func genPredictTOML() {
	item1 := &entity.RankParamItem{
		Model:                   "modelTest",
		Truncate:                false,
		FeatureAdditionalSchema: "set1",
	}
	item2 := &entity.RankParamItem{
		Model:                   "modelTest",
		Truncate:                true,
		FeatureAdditionalSchema: "set2",
	}
	item3 := &entity.RankParamItem{
		Model:                   "modelTest",
		Truncate:                true,
		FeatureAdditionalSchema: "set3",
	}
	item4 := &entity.RankParamItem{
		Model:                   "modelTest",
		Truncate:                true,
		FeatureAdditionalSchema: "set4",
	}
	params1 := make(map[string]*entity.RankParamItem)
	params1["sg"] = item1
	params1["id"] = item2
	params2 := make(map[string]*entity.RankParamItem)
	params2["vn"] = item3
	params2["my"] = item4
	rankParams1 := &entity.RankParams{Params: params1}
	rankParams2 := &entity.RankParams{Params: params2}
	c := &entity.PredictConfig{RankParams: map[string]*entity.RankParams{"sip": rankParams1, "intent_shopee": rankParams2}}
	// c := &entity.Config{PredictConfig: p}
	buffer := new(bytes.Buffer)
	encoder := toml.NewEncoder(buffer)
	if err := encoder.Encode(c); err != nil {
		panic(err)
	}
	fmt.Println("toml:")
	fmt.Println(buffer.String())
}

func main() {
	genPredictTOML()
}
