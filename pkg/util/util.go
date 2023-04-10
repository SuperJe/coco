package util

import (
	"bytes"
	"encoding/gob"
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
