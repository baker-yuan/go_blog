package schema_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Plugin List", func() {
	DescribeTable("test plugin basic",
		func(testCase base.HttpTestCase) {
			base.RunTestCase(testCase)
		},
		Entry("get all plugins", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/plugins",
			Query:        "all=true",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   []string{"request-id", "syslog", "echo", "proxy-mirror"},
			Sleep:        base.SleepTime,
		}),
		Entry("get all plugins", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/plugins",
			Query:        "all=false",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   []string{"request-id", "syslog", "echo", "proxy-mirror"},
			Sleep:        base.SleepTime,
		}),
	)
})
