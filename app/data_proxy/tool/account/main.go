package main

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	mongo2 "go.mongodb.org/mongo-driver/mongo"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/app/data_proxy/store"
	"github.com/SuperJe/coco/pkg/mongo"
)

var (
	prod int
	file string
)

func init() {
	flag.IntVar(&prod, "prod", 0, "1-request product address")
	flag.StringVar(&file, "file", "", "target file path")
}

func hashPassword(password string) string {
	password = strings.ToLower(password)
	// 添加盐值 "pepper"
	salt := "pepper" + password
	// 创建 SHA-512 哈希对象
	hash := sha512.New()
	// 计算哈希值
	hash.Write([]byte(salt))
	hashed := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashedStr := hex.EncodeToString(hashed)
	return hashedStr
}

// 每次创建都新建mongo和mysql实例, 数据量小, 不用在意此开销。
func registerUser(st *store.Student) error {
	if err := syncToMongo(st); err != nil {
		return err
	}
	return syncToAdmin(st)
}

func syncToMongo(st *store.Student) error {
	cli, err := mongo.NewCocoClient2()
	if err != nil {
		return err
	}
	user, err := cli.GetUserByName(context.Background(), st.Name)
	if err != nil && !errors.Is(err, mongo2.ErrNoDocuments) {
		return err
	}
	if user != nil {
		fmt.Println("====warning=====")
		fmt.Printf("%s has exist in cc\n\n", st.Name)
		return nil
	}
	return cli.RegisterUser(context.Background(), st.Name, st.Password)
}

func syncToAdmin(st *store.Student) error {
	data := st.ToReq()
	body, _ := json.Marshal(data)
	url := "http://127.0.0.1:7777/register"
	if prod == 1 {
		url = "http://81.71.3.223:7777/register"
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http err code:%d", rsp.StatusCode)
	}
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return errors.Wrap(err, "ReadAll err")
	}
	dataRsp := &model.RegisterRsp{}
	if err := json.Unmarshal(bs, dataRsp); err != nil {
		return errors.Wrap(err, "json unmarshal err")
	}
	if dataRsp.Code != 0 {
		return fmt.Errorf("register err:%s", dataRsp.Msg)
	}

	fmt.Printf("register %s success\n", st.Name)
	return nil
}

func main() {
	flag.Parse()
	// 读取整个文件内容
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	// 将内容按行分割
	lines := strings.Split(string(content), "\n")
	lines = lines[1:]
	for i, line := range lines {
		students := strings.Split(line, ";")
		if len(students) != 6 {
			fmt.Printf("line %d format err:%s, len:%d", i+1, line, len(students))
			return
		}
		st := &store.Student{Name: students[0], Password: students[1], Phone: students[2],
			Sex: students[3], Class: students[4], TeacherName: students[5]}
		if err := st.Valid(); err != nil {
			fmt.Printf("line %d invalid:%s", i+1, line)
			panic(err)
		}
		if err := registerUser(st); err != nil {
			panic(err)
		}
	}
}
