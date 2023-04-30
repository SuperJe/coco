package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/mongo/entity"
)

func update(user *entity.User, campProgression *model.CampaignProgression) error {
	data := &model.UpdateUserProgressionReq{
		Name:            user.Name,
		Completed:       user.CompletedLevelCount(),
		Gems:            user.GemCount(),
		LastLevel:       user.LastLevel,
		CampProgression: campProgression,
	}
	body, _ := json.Marshal(data)
	// req, err := http.NewRequest("POST", "81.71.3.223:9090", bytes.NewBuffer(body))
	req, err := http.NewRequest("POST", "http://127.0.0.1:7777/user_progression", bytes.NewBuffer(body))
	if err != nil {
		return errors.Wrap(err, "http.NewRequest err")
	}
	cli := &http.Client{}
	rsp, err := cli.Do(req)
	if err != nil {
		return errors.Wrap(err, "cli.Do err")
	}
	defer func() {
		if err := rsp.Body.Close(); err != nil {
			_ = rsp.Body.Close()
		}
	}()
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return errors.Wrap(err, "ReadAll err")
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http err code:%d", rsp.StatusCode)
	}

	fmt.Println("rsp:", string(bs))
	return nil
}

func get(name string) error {
	req, err := http.NewRequest("GET", "http://127.0.0.1:7777/user_progression", nil)
	if err != nil {
		return errors.Wrap(err, "http.NewRequest err")
	}
	params := req.URL.Query()
	params.Add("name", name)
	req.URL.RawQuery = params.Encode()
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "cli.Do err")
	}
	defer func() {
		if err := rsp.Body.Close(); err != nil {
			_ = rsp.Body.Close()
		}
	}()
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return errors.Wrap(err, "ReadAll err")
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http err code:%d", rsp.StatusCode)
	}
	data := &model.GetUserProgressionRsp{}
	if err := json.Unmarshal(bs, data); err != nil {
		return errors.Wrap(err, "unmarshal err")
	}
	fmt.Println("data:", data)
	return nil
}

func getCampProgression(ctx context.Context, name string) (*model.CampaignProgression, error) {
	cli, err := mongo.NewCocoClient2()
	if err != nil {
		return nil, errors.Wrap(err, "NewCocoClient2 err")
	}
	counts, err := cli.CountLevels(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cli.CountLevels err")
	}
	levels, err := cli.GetCompletedLevels(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "cli.GetCompletedLevels err")
	}
	completed, err := cli.GroupLevelByCampaign(ctx, levels)
	if err != nil {
		return nil, errors.Wrap(err, "cli.GroupLevelByCampaign err")
	}

	progressions := &model.CampaignProgression{}
	progressions.Dungeon = buildProgression("Dungeon", completed, counts)
	progressions.Forest = buildProgression("Forest", completed, counts)
	progressions.Desert = buildProgression("Desert", completed, counts)
	progressions.Mountain = buildProgression("Mountain", completed, counts)
	progressions.Glacier = buildProgression("Glacier", completed, counts)
	return progressions, nil
}

func buildProgression(campaign string, completed map[string][]string, counts map[string]int32) *model.Progression {
	return &model.Progression{
		Done:       int32(len(completed[campaign])),
		Unfinished: counts[campaign] - int32(len(completed[campaign])),
		Total:      counts[campaign],
	}
}

func main() {
	// TODO: 修改端口
	cache := make(map[string]*entity.User, 100)
	mgo, err := mongo.NewCocoClient2()
	if err != nil {
		panic(err)
	}
	for {
		users, err := mgo.GetUsers(context.Background())
		if err != nil {
			panic(err)
		}
		for _, user := range users {
			if len(user.Name) == 0 {
				continue
			}
			last, ok := cache[user.Name]
			if ok && last.LastLevel == user.LastLevel {
				fmt.Printf("user:%s stay the same level:%s\n", user.Name, user.LastLevel)
				continue
			}
			campProgression, err := getCampProgression(context.Background(), user.Name)
			if err != nil {
				fmt.Println("getCampProgression err:", err.Error())
				continue
			}
			if err := update(user, campProgression); err != nil {
				fmt.Println("doRequest err:", err.Error())
				continue
			}
			if err := get(user.Name); err != nil {
				fmt.Println("get err:", err.Error())
				continue
			}
			cache[user.Name] = user.Clone()
		}
		time.Sleep(3 * time.Second)
	}
}
