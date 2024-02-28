package follow

import (
	"github.com/baker-yuan/go-blog/interaction/application/dto/follow/command"
	dto "github.com/baker-yuan/go-blog/interaction/ui/http/dto/follow"
)

// GenAddFollowCMD 新增关注关系
func GenAddFollowCMD(req *dto.AddFollowReq) (cmd *command.AddFollowCMD) {
	cmd = &command.AddFollowCMD{}
	cmd.UID = req.UID
	cmd.FollowUID = req.FollowUID
	return cmd
}
