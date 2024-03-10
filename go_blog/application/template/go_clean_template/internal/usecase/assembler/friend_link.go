package assembler

import (
	pb "github.com/baker-yuan/go-blog/protocol/template"
	"github.com/baker-yuan/go-blog/template/go_clean_template/internal/entity"
	"google.golang.org/protobuf/proto"
)

// FriendLinkEntityToModel entity转pb
func FriendLinkEntityToModel(friendLink *entity.FriendLink) *pb.FriendLink {
	modelRes := &pb.FriendLink{
		Id: proto.Uint32(friendLink.ID),
	}
	return modelRes
}

// AddOrUpdateFriendLinkReqToEntity pb转entity
func AddOrUpdateFriendLinkReqToEntity(pbFriendLink *pb.AddOrUpdateFriendLinkReq) *entity.FriendLink {
	entityRes := &entity.FriendLink{}
	return entityRes
}
