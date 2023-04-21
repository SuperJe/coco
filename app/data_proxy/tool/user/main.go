package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/mongo/entity"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

func doRequest(user *entity.User) error {
	body, _ := json.Marshal(user)
	// req, err := http.NewRequest("POST", "81.71.3.223:9090", bytes.NewBuffer(body))
	req, err := http.NewRequest("POST", "http://127.0.0.1:9090/user_progression", bytes.NewBuffer(body))
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

	fmt.Println("bs:", string(bs))
	return nil
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
			if err := doRequest(user); err != nil {
				fmt.Println("doRequest err:", err.Error())
				continue
			}
			cache[user.Name] = user.Clone()
		}
		time.Sleep(3 * time.Second)
	}
}
