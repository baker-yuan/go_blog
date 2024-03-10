package conversion

import "strconv"

// ConversionUtils 类型转换结构定义
type ConversionUtils struct {
}

// StrToInt 字符串转int类型，当存在defaultNum时，出现异常会返回设置的默认值，若不出现异常正常返回
func (c ConversionUtils) StrToInt(strNum string, defaultNum ...int) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return num
}
