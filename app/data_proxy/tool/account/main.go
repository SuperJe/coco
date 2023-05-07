package main

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/SuperJe/coco/pkg/mongo"
	"github.com/SuperJe/coco/pkg/util/encode"
	"strings"
)

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

func registerUser(name, pwd string) {
	cli, err := mongo.NewCocoClient2()
	if err != nil {
		panic(err)
	}
	if err := cli.RegisterUser(context.Background(), name, pwd); err != nil {
		panic(err)
	}
	fmt.Println("NewUser success")
}

func main() {
	fmt.Println("pwd:", encode.Sha512WithSalt("jelly003", "pepper"))
	// registerUser("jelly001", "jelly001")
}
