package cqe

// PageResult 分页对象
type PageResult struct {
	RecordList interface{} `json:"recordList"` // 分页列表(集合)
	Count      uint32      `json:"count"`      // 总数
}
