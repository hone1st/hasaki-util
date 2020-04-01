package string

import (
	"runtime"
	"strings"
)

// 系统拼接符
// b为true则在前面加一个系统切割符
func SystemAppend(b bool, params ...string) string {
	var s = "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	if b {
		return s + strings.Join(params, s)
	}
	return strings.Join(params, s)
}

// 系统切割符
func SystemSplit(str string) []string {
	var s = "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	return strings.Split(str, s)
}

// 替换字符串。多种字符替换成一种字符
// 如果被替换的字符长度为空返回本身
func Replace(need string, new string, params ...string) string {
	if len(params) == 0 {
		return need
	}
	for _, v := range params {
		need = strings.Replace(need, v, new, -1)
	}
	return need
}

// 检测是否包含某个或者多个字符
// b为true时 被包含的目标需要全部存在才返回true
// b为false  只包含某个目标则返回true
// 如果params长度为0时返回false
func Contains(need string, b bool, params ...string) bool {
	if len(params) == 0 {
		return false
	}
	for _, v := range params {
		if (strings.Contains(need, v) && !b) || (!strings.Contains(need, v) && b) {
			return !b
		}
	}
	return b
}
