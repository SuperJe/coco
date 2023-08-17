package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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
