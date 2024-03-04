package slice

// FieldExtractor 是一个泛型结构体，提供了从切片中提取字段的方法。
// T 数组原始类型
// K 字段类型
type FieldExtractor[T any, K comparable] struct{}

// ExtractField 方法从切片中提取指定字段。
func (fe FieldExtractor[T, K]) ExtractField(slice []T, extractFunc func(T) K) []K {
	fieldSlice := make([]K, 0, len(slice)) // 预分配足够的空间
	for _, item := range slice {
		fieldSlice = append(fieldSlice, extractFunc(item))
	}
	return fieldSlice
}

// GroupByField 方法根据提取函数对切片中的元素进行分组。
// 它返回一个map，其中键是提取的字段，值是具有相同字段值的元素的切片。
func (fe FieldExtractor[T, K]) GroupByField(slice []T, extractFunc func(T) K) map[K][]T {
	groups := make(map[K][]T)
	for _, item := range slice {
		key := extractFunc(item)
		groups[key] = append(groups[key], item)
	}
	return groups
}
