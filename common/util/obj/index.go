package obj

import "encoding/json"

// ObjectUtils 对象工具
type ObjectUtils[T any] struct{}

// Unmarshal 反序列化
func (s *ObjectUtils[T]) Unmarshal(data string) (T, error) {
	var t T
	return t, json.Unmarshal([]byte(data), &t)
}
