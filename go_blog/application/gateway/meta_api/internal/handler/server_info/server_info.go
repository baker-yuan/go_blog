package server_info

import (
	"reflect"
	"strings"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/entity"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/store"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/shiningrush/droplet"
	"github.com/shiningrush/droplet/wrapper"
	wgin "github.com/shiningrush/droplet/wrapper/gin"
)

type Handler struct {
	serverInfoStore store.Interface
}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{
		serverInfoStore: store.GetStore(store.HubKeyServerInfo),
	}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/apisix/admin/server_info/:id", wgin.Wraps(h.Get, wrapper.InputType(reflect.TypeOf(GetInput{}))))
	r.GET("/apisix/admin/server_info", wgin.Wraps(h.List, wrapper.InputType(reflect.TypeOf(ListInput{}))))
}

type GetInput struct {
	ID string `auto_read:"id,path" validate:"required"`
}

func (h *Handler) Get(c droplet.Context) (interface{}, error) {
	input := c.Input().(*GetInput)

	r, err := h.serverInfoStore.Get(c.Context(), input.ID)
	if err != nil {
		return handler.SpecCodeResponse(err), err
	}

	return r, nil
}

type ListInput struct {
	store.Pagination
	Hostname string `auto_read:"hostname,query"`
}

func (h *Handler) List(c droplet.Context) (interface{}, error) {
	input := c.Input().(*ListInput)

	ret, err := h.serverInfoStore.List(c.Context(), store.ListInput{
		Predicate: func(obj interface{}) bool {
			if input.Hostname != "" {
				return strings.Contains(obj.(*entity.ServerInfo).Hostname, input.Hostname)
			}
			return true
		},
		PageSize:   input.PageSize,
		PageNumber: input.PageNumber,
	})

	if err != nil {
		return nil, err
	}

	return ret, nil
}
