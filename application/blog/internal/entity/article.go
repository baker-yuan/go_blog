package entity

import pb "github.com/baker-yuan/go-blog/protocol/blog"

// Article 文章表
type Article struct {
	// 基本数据
	ID          uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	CategoryID  uint32 `gorm:"index:idx_category_id;column:category_id;type:int(10);unsigned;not null;default:0;comment:文章分类表ID|biz_type表主键"`
	Title       string `gorm:"column:title;type:varchar(255);not null;default:'';comment:文章标题"`
	Description string `gorm:"column:description;type:varchar(200);not null;default:'';comment:文章简介，最多200字"`
	Content     string `gorm:"column:content;type:mediumtext;not null;comment:文章内容"`
	CoverImage  string `gorm:"column:cover_image;type:varchar(255);not null;default:'';comment:文章封面图片"`
	OriginalURL string `gorm:"column:original_url;type:varchar(255);not null;default:'';comment:原文链接"`
	//
	Password string `gorm:"column:password;type:varchar(255);not null;default:'';comment:密码保护"`
	// 辅助
	Words    uint32 `gorm:"column:words;type:int unsigned;not null;default:0;comment:文章字数"`
	ReadTime uint32 `gorm:"column:read_time;type:int unsigned;not null;default:0;comment:文章阅读时长(分钟)"`
	// 类型描述
	Type       pb.ArticleType       `gorm:"column:type;type:tinyint unsigned;not null;default:0;comment:文章类型 1-原创 2-转载 3-翻译"`
	Status     pb.ArticleStatus     `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:文章状态 0-草稿 1-已发布"`
	Format     pb.ArticleFormat     `gorm:"column:format;type:tinyint unsigned;not null;default:0;comment:文章格式 1-markdown 2-富文本"`
	Visibility pb.ArticleVisibility `gorm:"column:visibility;type:tinyint unsigned;not null;default:0;comment:文章可见性 1-公开 2-私密 3-密码保护"`
	// 标志
	IsTop            BoolBit `gorm:"column:is_top;type:bit(1);not null;default:b'0';comment:置顶开关"`
	IsRecommend      BoolBit `gorm:"column:is_recommend;type:bit(1);not null;default:b'0';comment:推荐开关"`
	IsAppreciation   BoolBit `gorm:"column:is_appreciation;type:bit(1);not null;default:b'0';comment:赞赏开关"`
	IsCommentEnabled BoolBit `gorm:"column:is_comment_enabled;type:bit(1);not null;default:b'0';comment:评论开关"`
	// 公共字段
	IsDeleted  BoolBit   `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否删除"`
	CreateTime Timestamp `gorm:"index:idx_create_time;column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime Timestamp `gorm:"index:idx_update_time;column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

func (Article) TableName() string {
	return "blog_article"
}

var (
	ArticleTbName                = "blog_article"
	ArticleFieldID               = "id"
	ArticleFieldCategoryID       = "category_id"
	ArticleFieldCoverImage       = "cover_image"
	ArticleFieldTitle            = "title"
	ArticleFieldContent          = "content"
	ArticleFieldVisibility       = "visibility"
	ArticleFieldPassword         = "password"
	ArticleFieldDescription      = "description"
	ArticleFieldType             = "type"
	ArticleFieldStatus           = "status"
	ArticleFieldFormat           = "format"
	ArticleFieldOriginalUrl      = "original_url"
	ArticleFieldWords            = "words"
	ArticleFieldReadTime         = "read_time"
	ArticleFieldIsTop            = "is_top"
	ArticleFieldIsRecommend      = "is_recommend"
	ArticleFieldIsAppreciation   = "is_appreciation"
	ArticleFieldIsCommentEnabled = "is_comment_enabled"
	ArticleFieldIsDeleted        = "is_deleted"
	ArticleFieldCreateTime       = "create_time"
	ArticleFieldUpdateTime       = "update_time"
)

type Articles []*Article
