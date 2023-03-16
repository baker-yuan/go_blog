package dto

// ArchiveDTO 归档文章
type ArchiveDTO struct {
	ID           uint32 `json:"id"`           // id
	ArticleTitle string `json:"articleTitle"` // 标题
	CreateTime   uint32 `json:"createTime"`   // 发表时间
}

// ArticleBackDTO 后台文章
type ArticleBackDTO struct {
	ID           uint32   `json:"id"`           // id
	ArticleCover string   `json:"articleCover"` // 文章封面
	ArticleTitle string   `json:"articleTitle"` // 标题
	CreateTime   uint32   `json:"createTime"`   // 发表时间
	LikeCount    uint32   `json:"likeCount"`    // 点赞量
	ViewsCount   uint32   `json:"viewsCount"`   // 浏览量
	CategoryName string   `json:"categoryName"` // 文章分类名
	TagDTOList   []TagDTO `json:"tagDTOList"`   // 文章标签
	Type         uint32   `json:"type"`         // 文章类型
	IsTop        uint32   `json:"isTop"`        // 是否置顶
	IsDelete     uint32   `json:"isDelete"`     // 是否删除
	Status       uint32   `json:"status"`       // 文章状态
}

// ArticleHomeDTO 首页文章
type ArticleHomeDTO struct {
	ID             uint32   `json:"id"`             // id
	ArticleCover   string   `json:"articleCover"`   // 文章缩略图
	ArticleTitle   string   `json:"articleTitle"`   // 标题
	ArticleContent string   `json:"articleContent"` // 内容
	CreateTime     uint32   `json:"createTime"`     // 发表时间
	IsTop          uint32   `json:"isTop"`          // 是否置顶
	Type           uint32   `json:"type"`           // 文章类型
	CategoryId     uint32   `json:"categoryId"`     // 文章分类id
	CategoryName   string   `json:"categoryName"`   // 文章分类名
	TagDTOList     []TagDTO `json:"tagDTOList"`     // 文章标签
}

// ArticleDTO 文章
type ArticleDTO struct {
	ID                   uint32                `json:"id"`                   // id
	ArticleCover         string                `json:"articleCover"`         // 文章缩略图
	ArticleTitle         string                `json:"articleTitle"`         // 标题
	ArticleContent       string                `json:"articleContent"`       // 内容
	LikeCount            uint32                `json:"likeCount"`            // 点赞量
	ViewsCount           uint32                `json:"viewsCount"`           // 浏览量
	Type                 uint32                `json:"type"`                 // 文章类型
	OriginalUrl          string                `json:"originalUrl"`          // 原文链接
	CreateTime           uint32                `json:"createTime"`           // 发表时间
	UpdateTime           uint32                `json:"updateTime"`           // 更新时间
	CategoryId           uint32                `json:"categoryId"`           // 文章分类id
	CategoryName         string                `json:"categoryName"`         // 文章分类名
	TagDTOList           []TagDTO              `json:"tagDTOList"`           // 文章标签
	LastArticle          ArticlePaginationDTO  `json:"lastArticle"`          // 上一篇文章
	NextArticle          ArticlePaginationDTO  `json:"nextArticle"`          // 下一篇文章
	RecommendArticleList []ArticleRecommendDTO `json:"recommendArticleList"` // 推荐文章列表
	NewestArticleList    []ArticleRecommendDTO `json:"newestArticleList"`    // 最新文章列表
}

// ArticlePaginationDTO 文章上下篇
type ArticlePaginationDTO struct {
	ID           uint32 `json:"id"`           // id
	ArticleCover string `json:"articleCover"` // 文章缩略图
	ArticleTitle string `json:"articleTitle"` // 标题
}

// ArticleRecommendDTO 推荐文章
type ArticleRecommendDTO struct {
	ID           uint32 `json:"id"`           // id
	ArticleCover string `json:"articleCover"` // 文章缩略图
	ArticleTitle string `json:"articleTitle"` // 标题
	CreateTime   uint32 `json:"createTime"`   // 创建时间
}
