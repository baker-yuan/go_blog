package assembler

import (
	"github.com/baker-yuan/go-blog/all_packaged_library/constant"
	pb "github.com/baker-yuan/go-blog/proto/user"
	"github.com/baker-yuan/go-blog/user/application/dto/query"
)

func GenGetUserListQRY(req *pb.AdminListUserReq) *query.GetUserListQRY {
	qry := &query.GetUserListQRY{}
	qry.LoginType = req.LoginType
	qry.Nickname = req.Nickname
	if req.Current != 0 {
		qry.Current = req.Current
	} else {
		qry.Current = constant.DEFAULT_CURRENT
	}
	if req.Size != 0 {
		qry.Size = req.Size
	} else {
		qry.Size = constant.DEFAULT_SIZE
	}

	return qry
}
