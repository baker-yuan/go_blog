package entity

// Article 文章
type Article struct {
	ID             uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime     uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime     uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	UserId         uint32 `json:"userId"`                           // 作者
	ArticleCover   string `json:"articleCover"`                     // 文章缩略图
	ArticleTitle   string `json:"articleTitle"`                     // 标题
	ArticleContent string `json:"articleContent"`                   // 内容
	Type           uint32 `json:"type"`                             // 文章类型
	OriginalUrl    string `json:"originalUrl"`                      // 原文链接
	IsTop          bool   `json:"isTop"`                            // 是否置顶
	IsDelete       bool   `json:"isDelete"`                         // 是否删除
	Status         uint8  `json:"status"`                           // 文章状态 1-公开 2-私密 3-评论可见
	CategoryId     uint32 `json:"categoryId"`                       // 文章分类
	Category       Category
	ArticleTag     []ArticleTag
}

// ArticleTag 标签
type ArticleTag struct {
	ID         uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	ArticleId  uint32 `json:"articleId"`                        // 文章id
	TagId      uint32 `json:"tagId"`                            // 标签id
}

// Category 分类
type Category struct {
	ID           uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime   uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime   uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	CategoryName string `json:"categoryName"`                     // 分类名
}

// Top 文章置顶
func (a *Article) Top() {
	a.IsTop = true
}

// CancelTop 取消置顶
func (a *Article) CancelTop() {
	a.IsTop = false
}
