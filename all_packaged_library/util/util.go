package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"strconv"
)

// StrToUInt32 String类型转uint32
// param strNum 需要转换的字符串
// param defaultNum 默认值
func StrToUInt32(strNum string, defaultNum ...uint32) uint32 {
	num, err := strconv.ParseInt(strNum, 10, 32)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return uint32(num)
}

// DeepCopyByGob 利用gob进行深拷贝
func DeepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}

// DeepCopyByJson 利用json进行深拷贝
func DeepCopyByJson(src, dst interface{}) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}

// Uint32 stores v in a new uint32 value and returns a pointer to it.
func Uint32(v uint32) *uint32 { return &v }
