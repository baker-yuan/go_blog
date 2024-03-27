package middlewares_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Gzip enable", func() {
	It("get index.html", func() {
		base.RunTestCase(base.HttpTestCase{
			Object:        base.ManagerApiExpect(),
			Method:        http.MethodGet,
			Path:          "/",
			Headers:       map[string]string{"Accept-Encoding": "gzip, deflate, br"},
			ExpectHeaders: map[string]string{"Content-Encoding": "gzip"},
		})
	})
})
