package middlewares_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Invalid Request", func() {
	It("double dot in URL path (arbitrary file index)", func() {
		base.RunTestCase(base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/../../../../etc/hosts",
			ExpectStatus: http.StatusForbidden,
		})
		base.RunTestCase(base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/.%2e/%2e%2e/../etc/hosts",
			ExpectStatus: http.StatusForbidden,
		})
	})
})
