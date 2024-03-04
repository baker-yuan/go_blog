package slice

// SliceUtils 切片常用工具集，是一个泛型结构体，提供了操作切片的泛型方法。
// T 是切片元素类型
type SliceUtils[T comparable] struct{}

// SliceExist 方法判断切片中是否存在某个元素。
func (s SliceUtils[T]) SliceExist(slice []T, value T) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

// SliceRemove 方法移除切片中的一个或多个元素。
func (s SliceUtils[T]) SliceRemove(slice []T, removes ...T) []T {
	res := make([]T, 0)
	for _, item := range slice {
		exist := false
		for _, remove := range removes {
			if item == remove {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, item)
		}
	}
	return res
}

// SliceUnique 方法去除切片中的重复元素，并且可以接受一个可选的过滤函数。
func (s SliceUtils[T]) SliceUnique(slice []T, filters ...func(T) bool) []T {
	uniqueMap := make(map[T]struct{})
	var result []T

	// 检查是否提供了过滤函数
	var filterFunc func(T) bool
	if len(filters) > 0 {
		filterFunc = filters[0]
	}

	for _, v := range slice {
		if _, exist := uniqueMap[v]; !exist {
			// 如果提供了过滤函数，并且当前元素满足过滤条件，则跳过
			if filterFunc != nil && filterFunc(v) {
				continue
			}
			uniqueMap[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// SliceFilter 根据传入的过滤函数过滤切片中的元素。
// filterFunc 是一个函数，用于决定是否过滤掉某个元素。如果 filterFunc 返回 true，则该元素被过滤掉。
func (s SliceUtils[T]) SliceFilter(slice []T, filterFunc func(T) bool) []T {
	var res []T
	for _, item := range slice {
		if !filterFunc(item) {
			res = append(res, item)
		}
	}
	return res
}

// IsSubset 检查 subset 中的所有字符串是否都包含在 superset 中
func (s SliceUtils[T]) IsSubset(superset []T, subset []T) bool {
	set := make(map[T]struct{})
	for _, item := range superset {
		set[item] = struct{}{}
	}

	for _, item := range subset {
		if _, exists := set[item]; !exists {
			return false
		}
	}

	return true
}
