package healthz

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/shiningrush/droplet"
	wgin "github.com/shiningrush/droplet/wrapper/gin"
)

type Handler struct {
}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/ping", wgin.Wraps(h.healthZHandler))
}

func (h *Handler) healthZHandler(c droplet.Context) (interface{}, error) {
	return "pong", nil
}
