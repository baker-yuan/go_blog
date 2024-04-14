package assembler

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// FriendLinkEntityListToModel entity转pb
func FriendLinkEntityListToModel(friendLinkList []*entity.FriendLink, uidToNameMap map[uint32]string) []*pb.FriendLink {
	res := make([]*pb.FriendLink, 0)
	for _, friendLink := range friendLinkList {
		res = append(res, FriendLinkEntityToModel(friendLink, uidToNameMap))
	}
	return res
}

// FriendLinkEntityToModel entity转pb
func FriendLinkEntityToModel(friendLink *entity.FriendLink, uidToNameMap map[uint32]string) *pb.FriendLink {
	modelRes := &pb.FriendLink{
		Id:           friendLink.ID,
		LinkName:     friendLink.LinkName,
		LinkAvatar:   friendLink.LinkAvatar,
		LinkAddress:  friendLink.LinkAddress,
		LinkIntro:    friendLink.LinkIntro,
		Status:       friendLink.Status,
		Sort:         friendLink.Sort,
		IsDeleted:    bool(friendLink.IsDeleted),
		CreateUserId: friendLink.CreateUserID,
		UpdateUserId: friendLink.UpdateUserID,
		CreateTime:   uint32(friendLink.CreateTime),
		UpdateTime:   uint32(friendLink.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[friendLink.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[friendLink.UpdateUserID]
	return modelRes
}

// AddOrUpdateFriendLinkReqToEntity pb转entity
func AddOrUpdateFriendLinkReqToEntity(pbFriendLink *pb.AddOrUpdateFriendLinkReq) *entity.FriendLink {
	entityRes := &entity.FriendLink{
		ID:          pbFriendLink.Id,
		LinkName:    pbFriendLink.LinkName,
		LinkAvatar:  pbFriendLink.LinkAvatar,
		LinkAddress: pbFriendLink.LinkAddress,
		LinkIntro:   pbFriendLink.LinkIntro,
		Status:      pbFriendLink.Status,
		Sort:        pbFriendLink.Sort,
	}
	return entityRes
}
