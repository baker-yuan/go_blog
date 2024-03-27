package healthz_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Healthy check", func() {
	It("ping manager-api", func() {
		base.RunTestCase(base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/ping",
			ExpectStatus: http.StatusOK,
			ExpectBody:   "pong",
			Sleep:        base.SleepTime,
		})
	})
})
