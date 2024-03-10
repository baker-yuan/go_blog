package usecase

import (
	"context"

	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/application/blog/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
)

// ArticleUseCase 文章管理
type ArticleUseCase struct {
	ICommonUseCase
	articleRepo IArticleRepo
}

// NewArticleUseCase 创建文章管理service
func NewArticleUseCase(
	commonUseCase ICommonUseCase,
	articleRepo IArticleRepo,
) *ArticleUseCase {
	return &ArticleUseCase{
		ICommonUseCase: commonUseCase,
		articleRepo:    articleRepo,
	}
}

// ArticleDetail 文章详情
func (c *ArticleUseCase) ArticleDetail(ctx context.Context, req *pb.ArticleDetailReq) (*pb.Article, error) {
	article, err := c.articleRepo.GetArticleByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbArticle := assembler.ArticleEntityToModel(article)
	return pbArticle, nil
}

// SearchArticle 文章搜索
func (c *ArticleUseCase) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) ([]*pb.Article, uint32, error) {
	articles, total, err := c.articleRepo.SearchArticle(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.Article, 0)
	for _, article := range articles {
		data = append(data, assembler.ArticleEntityToModel(article))
	}
	return data, total, nil
}

// AddOrUpdateArticle 添加修改文章
func (c *ArticleUseCase) AddOrUpdateArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (uint32, error) {

	if req.GetId() == 0 {
		return c.addArticle(ctx, req)
	} else {
		dbArticle, err := c.articleRepo.GetArticleByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateArticle(ctx, dbArticle, req)
	}
}

func (c *ArticleUseCase) getWordsAndReadTime() (uint32, uint32) {
	return 0, 0
}

func (c *ArticleUseCase) addArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (uint32, error) {
	article := assembler.AddOrUpdateArticleReqToEntity(req)
	//article.CreateTime = entity.Timestamp(time.Now().Unix())
	//article.UpdateTime = entity.Timestamp(time.Now().Unix())

	article.Words, article.ReadTime = c.getWordsAndReadTime()

	lastInsertID, err := c.articleRepo.Save(ctx, article)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_ARTICLE,
		"{}", article,
		"新增文章",
	)
	return lastInsertID, nil
}

func (c *ArticleUseCase) updateArticle(ctx context.Context, dbArticle *entity.Article, req *pb.AddOrUpdateArticleReq) (uint32, error) {
	saveArticle := assembler.AddOrUpdateArticleReqToEntity(req)
	saveArticle.CreateTime = dbArticle.CreateTime
	//saveArticle.UpdateTime = entity.Timestamp(time.Now().Unix())

	saveArticle.Words, saveArticle.ReadTime = c.getWordsAndReadTime()

	if err := c.articleRepo.UpdateByID(ctx, saveArticle); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_ARTICLE,
		dbArticle, saveArticle,
		"全字段修改文章",
	)

	return req.GetId(), nil
}

// DeleteArticle 删除文章
func (c *ArticleUseCase) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) error {
	article, err := c.articleRepo.GetArticleByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.articleRepo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_ARTICLE,
		article, "{}",
		"删除文章",
	)
	return nil
}
