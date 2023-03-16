package do

import "github.com/baker-yuan/go-blog/all_packaged_library/common"

// Category 分类
type Category struct {
	common.Model
	CategoryName string `json:"categoryName"` // 分类名
}

func (a Category) TableName() string {
	return "tb_category"
}
