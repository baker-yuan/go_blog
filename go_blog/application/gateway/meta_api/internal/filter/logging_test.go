package filter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRequestLogHandler(t *testing.T) {
	r := gin.New()
	logger := log.GetLogger(log.AccessLog)
	r.Use(RequestLogHandler(logger))
	r.GET("/", func(c *gin.Context) {
	})

	w := performRequest(r, "GET", "/", nil)
	assert.Equal(t, 200, w.Code)
}
