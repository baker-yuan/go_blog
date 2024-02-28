package do

// Menu 菜单
type Menu struct {
	ID         uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	Name       string `json:"name"`                             // 菜单名
	Path       string `json:"path"`                             // 路径
	Component  string `json:"component"`                        // 组件
	Icon       string `json:"icon"`                             // icon
	OrderNum   uint32 `json:"orderNum"`                         // 排序
	ParentId   uint32 `json:"parentId"`                         // 父id
	IsHidden   uint32 `json:"isHidden"`                         // 是否隐藏
}
