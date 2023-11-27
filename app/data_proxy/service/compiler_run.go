package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/SuperJe/coco/app/data_proxy/model"
	"github.com/SuperJe/coco/pkg/common"
	"github.com/SuperJe/coco/pkg/util"
)

func (s *Service) CompilerRun(c *gin.Context) {
	req := &model.RunCompilerReq{}
	rsp := &model.RunCompilerRsp{}
	rsp.Msg = "成功"
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if util.EmptyS(req.Lang) || util.EmptyS(req.Code) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	compiler := newCompiler(req.Lang, req.Code, req.Input)
	res, err := compiler.Run()
	if err != nil {
		rsp.Msg = "失败\n" + res + "\n" + err.Error()
		rsp.Code = common.ErrCompile
		c.AbortWithStatusJSON(http.StatusOK, rsp)
		return
	}
	rsp.OutPut = res
	c.AbortWithStatusJSON(http.StatusOK, rsp)
}

func genFile(suffix string) string {
	path, _ := os.Getwd()
	return fmt.Sprintf(path+"/%d_%d.%s", time.Now().UnixMilli(), rand.Int63(), suffix)
}

func removeFile(file string) {
	if err := os.Remove(file); err != nil {
		fmt.Printf("remove %s err:%s", file, err.Error())
	}
}

func runWithTimeout(cmd *exec.Cmd, t int) (string, error) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd.Stdout = stdOut
	cmd.Stderr = stdErr
	if err := cmd.Start(); err != nil {
		return stdErr.String(), err
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				util.PrintGoroutineStack(err)
			}
		}()
		if err := cmd.Wait(); err != nil {
			fmt.Println("wait err:", err.Error())
		}
	}()
	sec := 0
	for {
		if cmd.ProcessState != nil {
			// 子进程结束
			break
		} else {
			// 子进程还在执行中
			time.Sleep(time.Second)
			sec++
			if sec >= t {
				if err := cmd.Process.Kill(); err != nil {
					fmt.Printf("timeout, kill %d err %s\n", cmd.Process.Pid, err.Error())
				}
				return "", fmt.Errorf("运行超时")
			}
		}
	}
	if cmd.ProcessState.Success() {
		return stdOut.String(), nil
	}
	return stdErr.String(), fmt.Errorf("运行失败")
}

func newCompiler(lang, code, input string) Compiler {
	switch lang {
	case model.LangCPP:
		return &CPPCompiler{code: code, input: input}
	case model.LangPY:
		return &PYCompiler{code: code, input: input}
	default:
		return nil
	}
}

type Compiler interface {
	Run() (string, error)
}

type CPPCompiler struct {
	code, input string
}

func (cc *CPPCompiler) Run() (string, error) {
	// 新建cpp文件
	var cppFile string
	var inputFile string
	var execFile string
	defer func() {
		removeFile(cppFile)
		removeFile(inputFile)
		removeFile(execFile)
	}()
	cppFile = genFile("cpp")
	content := []byte(cc.code)
	if err := ioutil.WriteFile(cppFile, content, 0644); err != nil {
		return "", errors.Wrapf(err, "write to %s err", cppFile)
	}
	// 新建保存输入参数的文件
	inputFile = genFile("txt")
	content = []byte(cc.input)
	if err := ioutil.WriteFile(inputFile, content, 0644); err != nil {
		return "", errors.Wrapf(err, "write to %s err", inputFile)
	}
	// 编译
	execFile = cppFile[:strings.LastIndex(cppFile, ".cpp")]
	cmd := exec.Command("g++", cppFile, "-std=c++11", "-o", execFile)
	stdErr := &bytes.Buffer{}
	stdOut := &bytes.Buffer{}
	cmd.Stderr = stdErr
	cmd.Stdout = stdOut
	if err := cmd.Run(); err != nil {
		return stdErr.String(), err
	}
	// 运行
	cmd = exec.Command("bash", "-c", execFile+" < "+inputFile)
	return runWithTimeout(cmd, 30)
}

type PYCompiler struct {
	code, input string
}

func (pyc *PYCompiler) Run() (string, error) {
	// 新建py文件
	var pyFile string
	var inputFile string
	defer func() {
		removeFile(pyFile)
		removeFile(inputFile)
	}()
	pyFile = genFile("py")
	content := []byte(pyc.code)
	if err := ioutil.WriteFile(pyFile, content, 0644); err != nil {
		return "", errors.Wrapf(err, "write to %s err", pyFile)
	}
	// 新建保存输入参数的文件
	inputFile = genFile("txt")
	content = []byte(pyc.input)
	if err := ioutil.WriteFile(inputFile, content, 0644); err != nil {
		return "", errors.Wrapf(err, "write to %s err", inputFile)
	}
	// 运行
	cmd := exec.Command("bash", "-c", "cat "+inputFile+" | "+"python3 "+pyFile)
	return runWithTimeout(cmd, 30)
}
