
// {{$.Name}}HttpSrv http服务
type {{$.Name}}HttpSrv struct {
	server {{ $.ServiceName }} // grpc生成的server
	router gin.IRouter // gin路由
}

// 客户端调用该方法注册
func Reg{{ $.ServiceName }}HttpSrv(srv {{ $.ServiceName }}, r ...gin.IRouter) {
    var engine gin.IRouter

    // 优先使用传入的router
    if len(r) != 0 {
        engine = r[0]
    } else {
        engine = gin.Default()
    }

	s := {{.Name}}HttpSrv{
		server: srv,
		router: engine,
	}
	s.RegSrv()
}

{{range .Methods}}
func (s *{{$.Name}}HttpSrv) {{ .HandlerName }} (c *gin.Context) {
	var in {{.Request}}
{{if eq .Method "GET" "DELETE" }}
	if err := c.ShouldBindQuery(&in); err != nil {
        rsp := Response{
            Code: http.StatusBadRequest,
            Message: "",
        }
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
{{else if eq .Method "POST" "PUT" }}
	if err := c.ShouldBindJSON(&in); err != nil {
        rsp := Response{
            Code: http.StatusBadRequest,
            Message: "",
        }
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
{{else}}
	if err := c.ShouldBind(&in); err != nil {
        rsp := Response{
            Code: http.StatusBadRequest,
            Message: "",
        }
		c.JSON(http.StatusBadRequest, rsp)
        return
	}
{{end}}
{{if .HasPathParams }}
	{{range $item := .PathParams}}
	in.{{$.GoCamelCase $item }} = c.Params.ByName("{{$item}}")
	{{end}}
{{end}}
	out, err := s.server.{{.Name}}(c, &in)
	if err != nil {
        rsp := Response{
            Code: http.StatusInternalServerError,
            Message: "",
        }
		c.JSON(http.StatusInternalServerError, rsp)
        return
	}
	c.JSON(http.StatusOK, out)
}
{{end}}

// RegSrv 路由和对应的处理方法进行绑定
func (s *{{$.Name}}HttpSrv) RegSrv() {
{{range .Methods}}
		s.router.Handle("{{.Method}}", "{{.Path}}", s.{{ .HandlerName }})
{{end}}
}