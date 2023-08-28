package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	path, _ := os.Getwd()
	path += "/app/data_proxy/tool/compiler/"
	cmd := exec.Command("g++", path+"main.cpp" /* "-std=c++11",*/, "-o", path+"main")
	stdErr := &bytes.Buffer{}
	stdOut := &bytes.Buffer{}
	cmd.Stderr = stdErr
	cmd.Stdout = stdOut
	if err := cmd.Run(); err != nil {
		fmt.Println("err:", err)
		fmt.Println("std err:", stdErr.String())
		return
	}
	// 执行
	cmd = exec.Command(path + "main")
	cmd.Stdout = stdOut
	cmd.Stderr = stdErr
	if err := cmd.Run(); err != nil {
		fmt.Println("err:", err)
		fmt.Println("std err:", stdErr.String())
		return
	}
	fmt.Printf("exec result:\n\t%s\n", stdOut.String())
}
