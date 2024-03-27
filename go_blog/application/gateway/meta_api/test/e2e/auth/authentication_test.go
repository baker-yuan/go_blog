package auth_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Authentication", func() {
	DescribeTable("test auth module",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("Access with valid authentication token", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/routes",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   `"code":0`,
		}),
		Entry("Access with malformed authentication token", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/routes",
			Headers:      map[string]string{"Authorization": "Not-A-Valid-Token"},
			ExpectStatus: http.StatusUnauthorized,
			ExpectBody:   `"message":"request unauthorized"`,
		}),
		Entry("Access without authentication token", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/routes",
			ExpectStatus: http.StatusUnauthorized,
			ExpectBody:   `"message":"request unauthorized"`,
		}),
	)
})
