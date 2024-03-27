package schema

import (
	"reflect"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/handler"
	"github.com/gin-gonic/gin"
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
	r.GET("/apisix/admin/plugins", wgin.Wraps(h.Plugins, wrapper.InputType(reflect.TypeOf(ListInput{}))))
}

type ListInput struct {
	All bool `auto_read:"all,query"`
}

func (h *Handler) Plugins(c droplet.Context) (interface{}, error) {
	input := c.Input().(*ListInput)

	plugins := conf.Schema.Get("plugins")
	if input.All {
		var res []map[string]interface{}
		list := plugins.Value().(map[string]interface{})
		for name, schemaConfig := range list {
			if enable, ok := conf.Plugins[name]; !ok || !enable {
				continue
			}
			plugin := schemaConfig.(map[string]interface{})
			plugin["name"] = name
			if _, ok := plugin["type"]; !ok {
				plugin["type"] = "other"
			}
			res = append(res, plugin)
		}
		return res, nil
	}

	var ret []string
	list := plugins.Map()
	for pluginName := range list {
		if enable, ok := conf.Plugins[pluginName]; !ok || !enable {
			continue
		}

		ret = append(ret, pluginName)
	}

	return ret, nil
}
