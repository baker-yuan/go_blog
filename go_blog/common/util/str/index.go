package str

import (
	"encoding/json"
	"strings"
)

// StrUtils 字符串处理
type StrUtils struct {
}

// Obj2Json 对象转JSON
func (c StrUtils) Obj2Json(obj interface{}) string {
	if obj == nil {
		return ""
	}
	data, _ := json.Marshal(obj)
	return string(data)
}

// RemoveRepeatAndBlank 切片元素去重，去除切片中的空字符串
func (c StrUtils) RemoveRepeatAndBlank(arr []string) []string {
	return removeRepeatAndBlank(arr)
}

func removeRepeatAndBlank(arr []string) []string {
	out := make([]string, 0)
	if len(arr) == 0 {
		return out
	}

	tmp := make(map[string]struct{})
	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		if _, exist := tmp[v]; !exist {
			tmp[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// RemoveBlank 去除空字符串
func (c StrUtils) RemoveBlank(arr []string) []string {
	out := make([]string, 0)
	if len(arr) == 0 {
		return out
	}

	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		out = append(out, v)
	}
	return out
}

// SplitString 根据传入的多个分隔符切割字符串
func (c StrUtils) SplitString(s string, separators ...rune) []string {
	// 定义一个分隔函数，如果当前字符是分隔符中的任意一个，则返回true
	separatorFunc := func(c rune) bool {
		for _, sep := range separators {
			if c == sep {
				return true
			}
		}
		return false
	}
	return removeRepeatAndBlank(strings.FieldsFunc(s, separatorFunc))
}
