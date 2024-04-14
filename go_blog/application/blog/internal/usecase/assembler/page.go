package assembler

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// PageEntityListToModel entity转pb
func PageEntityListToModel(pageList []*entity.Page, uidToNameMap map[uint32]string) []*pb.Page {
	res := make([]*pb.Page, 0)
	for _, page := range pageList {
		res = append(res, PageEntityToModel(page, uidToNameMap))
	}
	return res
}

// PageEntityToModel entity转pb
func PageEntityToModel(page *entity.Page, uidToNameMap map[uint32]string) *pb.Page {
	modelRes := &pb.Page{
		Id:           page.ID,
		PageName:     page.PageName,
		PageLabel:    page.PageLabel,
		PageCover:    page.PageCover,
		CreateUserId: page.CreateUserID,
		UpdateUserId: page.UpdateUserID,
		CreateTime:   uint32(page.CreateTime),
		UpdateTime:   uint32(page.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[page.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[page.UpdateUserID]
	return modelRes
}

// AddOrUpdatePageReqToEntity pb转entity
func AddOrUpdatePageReqToEntity(pbPage *pb.AddOrUpdatePageReq) *entity.Page {
	entityRes := &entity.Page{
		ID:        pbPage.Id,
		PageName:  pbPage.PageName,
		PageLabel: pbPage.PageLabel,
		PageCover: pbPage.PageCover,
	}
	return entityRes
}
