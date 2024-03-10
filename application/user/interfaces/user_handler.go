package interfaces

import (
	"context"

	"github.com/baker-yuan/go-blog/application/user/application"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/auth"
	pb "github.com/baker-yuan/go-blog/protocol/user"
	"github.com/gin-gonic/gin"
)

// 强制Users实现UserApiService
var _ pb.UserApiService = &Users{}

// Users struct defines the dependencies that will be used
type Users struct {
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

func (s *Users) SearchUser(ctx context.Context, req *pb.SearchUserReq) (*pb.SearchUserRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Users) AddOrUpdateUser(ctx context.Context, req *pb.AddOrUpdateUserReq) (*pb.AddOrUpdateRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Users) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Users) UserDetail(ctx context.Context, req *pb.UserDetailReq) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

// NewUsers Users constructor
func NewUsers(us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Users {
	return &Users{
		us: us,
		rd: rd,
		tk: tk,
	}
}

func (s *Users) SaveUser(c *gin.Context) {
	//var user entity.User
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"invalid_json": "invalid json",
	//	})
	//	return
	//}
	//// validate the request:
	//validateErr := user.Validate("")
	//if len(validateErr) > 0 {
	//	c.JSON(http.StatusUnprocessableEntity, validateErr)
	//	return
	//}
	//newUser, err := s.us.SaveUser(&user)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err)
	//	return
	//}
	//c.JSON(http.StatusCreated, newUser.PublicUser())
}

func (s *Users) GetUsers(c *gin.Context) {
	//users := entity.Users{} // customize user
	//var err error
	//// us, err = application.UserApp.GetUsers()
	//users, err = s.us.GetUsers()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//c.JSON(http.StatusOK, users.PublicUsers())
}

func (s *Users) GetUser(c *gin.Context) {
	//userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//user, err := s.us.GetUser(userId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//c.JSON(http.StatusOK, user.PublicUser())
}
