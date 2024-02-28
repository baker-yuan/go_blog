package po

import "github.com/baker-yuan/go-blog/all_packaged_library/do"

// CategoryPO 分类
type CategoryPO struct {
	do.Model
	CategoryName string `json:"categoryName"` // 分类名
}

func (a CategoryPO) TableName() string {
	return "tb_category"
}
