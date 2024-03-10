package assembler

import (
	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
)

// ArticleEntityToModel entity转pb
func ArticleEntityToModel(article *entity.Article) *pb.Article {
	modelRes := &pb.Article{
		// 基本数据
		Id:          article.ID,
		CategoryId:  article.CategoryID,
		Title:       article.Title,
		Description: article.Description,
		Content:     article.Content,
		CoverImage:  article.CoverImage,
		OriginalUrl: article.OriginalURL,
		// 辅助信息
		Password: article.Password,
		Words:    article.Words,
		ReadTime: article.ReadTime,
		// 类型描述
		Type:       article.Type,
		Status:     article.Status,
		Format:     article.Format,
		Visibility: article.Visibility,
		// 标志位
		IsTop:            bool(article.IsTop),
		IsRecommend:      bool(article.IsRecommend),
		IsAppreciation:   bool(article.IsAppreciation),
		IsCommentEnabled: bool(article.IsCommentEnabled),
		// 公共字段
		IsDeleted:  bool(article.IsDeleted),
		CreateTime: uint32(article.CreateTime),
		UpdateTime: uint32(article.UpdateTime),
	}
	return modelRes
}

// AddOrUpdateArticleReqToEntity pb转entity
func AddOrUpdateArticleReqToEntity(pbArticle *pb.AddOrUpdateArticleReq) *entity.Article {
	entityRes := &entity.Article{
		// 基本数据
		ID:          pbArticle.Id,
		CategoryID:  pbArticle.CategoryId,
		Title:       pbArticle.Title,
		Description: pbArticle.Description,
		Content:     pbArticle.Content,
		CoverImage:  pbArticle.CoverImage,
		OriginalURL: pbArticle.OriginalUrl,
		// 辅助信息
		Password: pbArticle.Password,
		// 类型描述
		Type:       pbArticle.Type,
		Status:     pbArticle.Status,
		Format:     pbArticle.Format,
		Visibility: pbArticle.Visibility,
		// 标志位
		IsTop:            entity.BoolBit(pbArticle.IsTop),
		IsRecommend:      entity.BoolBit(pbArticle.IsRecommend),
		IsAppreciation:   entity.BoolBit(pbArticle.IsAppreciation),
		IsCommentEnabled: entity.BoolBit(pbArticle.IsCommentEnabled),
	}
	return entityRes
}
