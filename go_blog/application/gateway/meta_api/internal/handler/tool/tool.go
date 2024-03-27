package tool

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/handler"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/shiningrush/droplet"
	wgin "github.com/shiningrush/droplet/wrapper/gin"
)

type Handler struct {
}

type InfoOutput struct {
	Hash    string `json:"commit_hash"`
	Version string `json:"version"`
}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/apisix/admin/tool/version", wgin.Wraps(h.Version))
}

func (h *Handler) Version(_ droplet.Context) (interface{}, error) {
	hash, version := utils.GetHashAndVersion()
	return &InfoOutput{
		Hash:    hash,
		Version: version,
	}, nil
}
