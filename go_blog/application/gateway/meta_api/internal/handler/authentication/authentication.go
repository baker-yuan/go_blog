package authentication

import (
	"reflect"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/handler"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils/consts"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/shiningrush/droplet"
	"github.com/shiningrush/droplet/wrapper"
	wgin "github.com/shiningrush/droplet/wrapper/gin"
)

type Handler struct {
}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.POST("/apisix/admin/user/login", wgin.Wraps(h.userLogin, wrapper.InputType(reflect.TypeOf(LoginInput{}))))
}

type UserSession struct {
	Token string `json:"token"`
}

// swagger:model LoginInput
type LoginInput struct {
	// user name
	Username string `json:"username" validate:"required"`
	// password
	Password string `json:"password" validate:"required"`
}

// swagger:operation POST /apisix/admin/user/login userLogin
//
// user login.
//
// ---
// produces:
// - application/json
// parameters:
//   - name: username
//     in: body
//     description: user name
//     required: true
//     type: string
//   - name: password
//     in: body
//     description: password
//     required: true
//     type: string
//
// responses:
//
//	'0':
//	  description: login success
//	  schema:
//	    "$ref": "#/definitions/ApiError"
//	default:
//	  description: unexpected error
//	  schema:
//	    "$ref": "#/definitions/ApiError"
func (h *Handler) userLogin(c droplet.Context) (interface{}, error) {
	input := c.Input().(*LoginInput)
	username := input.Username
	password := input.Password

	// /api/conf/conf.yaml#users
	user := conf.UserList[username]
	if username != user.Username || password != user.Password {
		return nil, consts.ErrUsernamePassword
	}

	// create JWT for session
	claims := jwt.StandardClaims{
		Subject:   username,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(conf.AuthConf.ExpireTime)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(conf.AuthConf.Secret))

	// output token
	return &UserSession{
		Token: signedToken,
	}, nil
}
