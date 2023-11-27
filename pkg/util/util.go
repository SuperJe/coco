package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
)

// DeepCopy 使用gob序列化, 默认不使用json是因为json不能序列化为未导出的数据
func DeepCopy(src, dst interface{}) error {
	var buffer *bytes.Buffer
	enc := gob.NewEncoder(buffer)
	dec := gob.NewDecoder(buffer)
	if err := enc.Encode(src); err != nil {
		return err
	}
	return dec.Decode(dst)
}

// JSONString 序列化为json, 忽略错误
func JSONString(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

// EmptyS 校验字符串是否为空
func EmptyS(s string) bool {
	return len(s) == 0
}

// ExistEmptyString 是否存在空字符串
func ExistEmptyString(s ...string) bool {
	for _, v := range s {
		if EmptyS(v) {
			return true
		}
	}
	return false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandString 高效产生n个随机字符串
func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func PrintGoroutineStack(i interface{}) {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("return value:%+v, %s\n", i, string(buf[:n]))
}
