package do

// Model 公共字段
type Model struct {
	ID         uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
}
