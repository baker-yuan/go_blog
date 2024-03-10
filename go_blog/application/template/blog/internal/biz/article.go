package biz

import (
	"context"
	"database/sql/driver"
	"errors"
	"time"

	pb "github.com/baker-yuan/go-blog/application/blog/api/blog/v1"
)

// Article 文章表
type Article struct {
	// 基本数据
	ID          uint32 `gorm:"column:id;type:int unsigned;primary_key;auto_increment;comment:主键"`
	CategoryID  uint32 `gorm:"column:category_id;type:int(10);unsigned;not null;default:0;comment:文章分类表ID|biz_type表主键"`
	Title       string `gorm:"column:title;type:varchar(255);not null;default:'';comment:文章标题"`
	Description string `gorm:"column:description;type:varchar(300);not null;default:'';comment:文章简介，最多200字"`
	Content     string `gorm:"column:content;type:mediumtext;not null;comment:文章内容"`
	CoverImage  string `gorm:"column:cover_image;type:varchar(255);not null;default:'';comment:文章封面图片"`
	OriginalURL string `gorm:"column:original_url;type:varchar(255);not null;default:'0';comment:原文链接"`
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
	CreateTime Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

func (Article) TableName() string {
	return "blog_article"
}

// Timestamp 用于处理数据库中 timestamp 类型和 Go 中 uint32 类型之间的转换
type Timestamp uint32

// Scan 实现 sql.Scanner 接口，用于从数据库读取值时的自定义扫描逻辑
func (ts *Timestamp) Scan(value interface{}) error {
	if value == nil {
		*ts = 0
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return errors.New("timestamp scan: type assertion to time.Time failed")
	}

	*ts = Timestamp(t.Unix())
	return nil
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的自定义值逻辑
func (ts Timestamp) Value() (driver.Value, error) {
	// 将 uint32 转换为 time.Time
	t := time.Unix(int64(ts), 0)
	return t, nil
}

// BoolBit 用于处理 MySQL 中 bit(1) 类型和 Go 中 bool 类型之间的转换
type BoolBit bool

// Scan 实现 sql.Scanner 接口，用于从数据库读取值时的自定义扫描逻辑
func (bb *BoolBit) Scan(value interface{}) error {
	if value == nil {
		*bb = false
		return nil
	}

	bv, ok := value.([]byte)
	if !ok {
		return errors.New("boolBit scan: type assertion to []byte failed")
	}

	*bb = bv[0] == 1
	return nil
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的自定义值逻辑
func (bb BoolBit) Value() (driver.Value, error) {
	if bb {
		return []byte{1}, nil
	}
	return []byte{0}, nil
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

type ArticleRepo interface {
	// GetArticleByID 根据文章id集合查询文章
	GetArticleByID(ctx context.Context, id int) (*Article, error)
	// GetArticleByIDs 根据文章id集合查询文章
	GetArticleByIDs(ctx context.Context, ids []int) (Articles, error)
	// Save 保存文章
	Save(ctx context.Context, article *Article) (uint32, error)
	// UpdateByID 根据ID修改文章
	UpdateByID(ctx context.Context, article *Article) error
	// DeleteByID 根据ID删除文章
	DeleteByID(ctx context.Context, id int) error
	// SearchArticle 文章搜索
	SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (Articles, uint32, error)
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

// ArticleEntityToModel entity转pb
func ArticleEntityToModel(article *Article) *pb.Article {
	modelRes := &pb.Article{
		Id:               article.ID,
		CategoryId:       article.CategoryID,
		CoverImage:       article.CoverImage,
		Title:            article.Title,
		Content:          article.Content,
		Visibility:       article.Visibility,
		Password:         article.Password,
		Description:      article.Description,
		Type:             article.Type,
		Status:           article.Status,
		Format:           article.Format,
		OriginalUrl:      article.OriginalURL,
		Words:            article.Words,
		ReadTime:         article.ReadTime,
		IsTop:            bool(article.IsTop),
		IsRecommend:      bool(article.IsRecommend),
		IsAppreciation:   bool(article.IsAppreciation),
		IsCommentEnabled: bool(article.IsCommentEnabled),
		IsDeleted:        bool(article.IsDeleted),
		CreateTime:       uint32(article.CreateTime),
		UpdateTime:       uint32(article.UpdateTime),
	}
	return modelRes
}

// AddOrUpdateArticleReqToEntity pb转entity
func AddOrUpdateArticleReqToEntity(pbArticle *pb.AddOrUpdateArticleReq) *Article {
	entityRes := &Article{}
	return entityRes
}

// SearchArticle 文章搜索
func (b *ArticleUsecase) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) ([]*pb.Article, uint32, error) {
	articles, total, err := b.repo.SearchArticle(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.Article, 0)
	for _, article := range articles {
		data = append(data, ArticleEntityToModel(article))
	}
	return data, total, nil
}

// AddOrUpdateArticle 添加修改文章
func (b *ArticleUsecase) AddOrUpdateArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (*pb.AddOrUpdateRsp, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteArticle 删除文章
func (b *ArticleUsecase) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

// ArticleDetail 文章详情
func (b *ArticleUsecase) ArticleDetail(ctx context.Context, req *pb.ArticleDetailReq) (*pb.Article, error) {
	//TODO implement me
	panic("implement me")
}
