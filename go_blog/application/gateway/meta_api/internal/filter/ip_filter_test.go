package filter

import (
	"net/http/httptest"
	"testing"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIPFilter_Handle(t *testing.T) {
	// empty allowed ip list --> should normal
	conf.AllowList = []string{}
	r := gin.New()
	r.Use(IPFilter())

	r.GET("/", func(c *gin.Context) {
	})

	w := performRequest(r, "GET", "/", nil)
	assert.Equal(t, 200, w.Code)

	// should forbidden
	conf.AllowList = []string{"10.0.0.0/8", "10.0.0.1"}
	r = gin.New()
	r.Use(IPFilter())
	r.GET("/fbd", func(c *gin.Context) {
	})

	w = performRequest(r, "GET", "/fbd", nil)
	assert.Equal(t, 403, w.Code)

	// should allowed
	conf.AllowList = []string{"10.0.0.0/8", "0.0.0.0/0"}
	r = gin.New()
	r.Use(IPFilter())
	r.GET("/test", func(c *gin.Context) {
	})
	w = performRequest(r, "GET", "/test", nil)
	assert.Equal(t, 200, w.Code)

	// should forbidden
	conf.AllowList = []string{"127.0.0.1"}
	r = gin.New()
	r.Use(IPFilter())
	r.GET("/test", func(c *gin.Context) {})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Forwarded-For", "127.0.0.1")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)

	req = httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Real-Ip", "127.0.0.1")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}
