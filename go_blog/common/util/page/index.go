package page

// PageUtils 分页工具结构体
type PageUtils[T any] struct{}

// Page 分页结果
type Page[T any] struct {
	Items      []T    // 当前页的项
	TotalItems uint32 // 总项数
	TotalPages uint32 // 总页数
	PageNum    uint32 // 当前页码
	PageSize   uint32 // 每页项数
}

// Paginate 对任意类型的切片进行分页
func (pu *PageUtils[T]) Paginate(items []T, pageNum uint32, pageSize uint32) Page[T] {
	totalItems := uint32(len(items))
	if totalItems == 0 || pageSize == 0 {
		// 如果没有项或pageSize为0，则返回空的分页结果
		return Page[T]{
			Items:      []T{},
			TotalItems: 0,
			TotalPages: 0,
			PageNum:    0,
			PageSize:   pageSize,
		}
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	// 确保页码在合理范围内
	if pageNum < 1 {
		pageNum = 1
	}
	if pageNum > totalPages {
		pageNum = totalPages
	}

	// 计算分页的起始和结束索引
	startIndex := (pageNum - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > totalItems {
		endIndex = totalItems
	}

	// 获取当前页的项
	pageItems := items[startIndex:endIndex]

	return Page[T]{
		Items:      pageItems,
		TotalItems: totalItems,
		TotalPages: totalPages,
		PageNum:    pageNum,
		PageSize:   pageSize,
	}
}
