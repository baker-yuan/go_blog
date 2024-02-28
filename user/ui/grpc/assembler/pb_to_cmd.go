package assembler

import (
	pb "github.com/baker-yuan/go-blog/proto/user"
	"github.com/baker-yuan/go-blog/user/application/dto/command"
)

// GenNamePassLoginCMD Generate Login CMD
func GenNamePassLoginCMD(req *pb.AdminLoginReq) (cmd *command.NamePassLoginCMD) {
	cmd = &command.NamePassLoginCMD{}
	cmd.Username = req.Username
	cmd.Password = req.Password
	return cmd
}

// GenChangePwdCMD Generate AdminUpdatePwd CMD
func GenChangePwdCMD(username string, req *pb.AdminUpdatePwdReq) (cmd *command.ChangePwdCMD) {
	cmd = &command.ChangePwdCMD{}
	cmd.OldPassword = req.OldPassword
	cmd.NewPassword = req.NewPassword
	cmd.Username = username
	return cmd
}
