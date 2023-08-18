package common

import (
	"fmt"

	"github.com/SuperJe/coco/pkg/util"
)

// CheckRegisterParam 检查注册参数
func CheckRegisterParam(name, pwd string) error {
	if util.EmptyS(name) || util.EmptyS(pwd) {
		return fmt.Errorf("invalid params, name:%s, pwd:%s", name, pwd)
	}
	if name[0] == ' ' || name[len(name)-1] == ' ' {
		return fmt.Errorf("user name can not begin or end with space")
	}
	// 用户名只能是空格, 数字和字母的组合, 且不能有连续的空格
	spacePos := -2
	for i, ch := range name {
		if ch == ' ' {
			if spacePos+1 == i {
				return fmt.Errorf("user name can not contain consecutive spaces")
			}
			spacePos = i
			continue
		}
		if ch >= '0' && ch <= '9' {
			continue
		}
		if ch >= 'a' && ch <= 'z' {
			continue
		}
		if ch >= 'A' && ch <= 'Z' {
			continue
		}
		return fmt.Errorf("unsupported character:%c", ch)
	}
	return nil
}
