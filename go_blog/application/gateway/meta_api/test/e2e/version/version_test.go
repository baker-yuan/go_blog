package version_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Version", func() {
	DescribeTable("version test",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("get version", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/tool/version",
			ExpectStatus: http.StatusOK,
			ExpectBody:   []string{"commit_hash", "\"version\""},
		}),
	)
})
