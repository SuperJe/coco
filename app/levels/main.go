package main

import (
	"context"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	s, err := newService()
	if err != nil {
		fmt.Printf("newService err:%s\n", err.Error())
		panic(err)
	}
	switch method {
	// 此选项不再维护, 之后不再从原关卡拷贝关卡再发布到地图, 而是自己编辑关卡之后发布到地图
	case "copy-level":
		result, err := s.copyLevel()
		if err != nil {
			fmt.Println("err:", err.Error())
			panic(err)
		}
		if err := s.insertLevelToCampaign(result.srcID, result.dstID); err != nil {
			fmt.Println("err:", err.Error())
			panic(err)
		}
	case "publish":
		if err := s.publish(context.Background()); err != nil {
			fmt.Println("err:", err.Error())
			panic(err)
		}
	case "print-config":
		c, err := configLoad(configPath)
		fmt.Printf("config:%+v, err:%+v\n", c, err)
	default:
		fmt.Println("invalid method")
	}
}
