package praise

import (
	"github.com/baker-yuan/go-blog/interaction/application/dto/praise/command"
	dto "github.com/baker-yuan/go-blog/interaction/ui/http/dto/praise"
)

// GenObjectPraiseCMD 新增点赞
func GenObjectPraiseCMD(req *dto.ObjectPraiseReq) (cmd *command.ObjectPraiseCMD) {
	cmd = &command.ObjectPraiseCMD{}
	cmd.ModuleCode = req.ModuleCode
	cmd.ObjectId = req.ObjectId
	cmd.Uid = req.Uid
	return cmd
}

func GenCancelObjectPraiseCMD(req *dto.CancelObjectPraiseReq) (cmd *command.CancelObjectPraiseCMD) {
	cmd = &command.CancelObjectPraiseCMD{}
	cmd.ModuleCode = req.ModuleCode
	cmd.ObjectId = req.ObjectId
	cmd.Uid = req.Uid
	return cmd
}
