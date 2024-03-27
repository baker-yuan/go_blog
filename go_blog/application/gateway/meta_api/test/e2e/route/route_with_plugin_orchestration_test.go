package route_test

import (
	"io/ioutil"
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
)

var _ = Describe("route with plugin orchestration", func() {
	bytes, err := ioutil.ReadFile("../../testdata/dag-conf.json")
	It("panics if readfile dag-conf.json error", func() {
		Expect(err).To(BeNil())
	})
	dagConf := string(bytes)

	// invalid dag config that not specified root node
	bytes, err = ioutil.ReadFile("../../testdata/invalid-dag-conf.json")
	It("panics if readfile invalid-dag-conf.json error", func() {
		Expect(err).To(BeNil())
	})
	invalidDagConf := string(bytes)

	DescribeTable("test route with plugin orchestration",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("make sure the route is not created", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusNotFound,
			ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
		}),
		Entry("create route with invalid dag config", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodPut,
			Path:         "/apisix/admin/routes/r1",
			Body:         invalidDagConf,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusBadRequest,
		}),
		Entry("make sure the route created failed", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusNotFound,
			ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
			Sleep:        base.SleepTime,
		}),
		Entry("create route with correct dag config", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodPut,
			Path:         "/apisix/admin/routes/r1",
			Body:         dagConf,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify the route(should be blocked)", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Query:        "t=root.exe",
			ExpectStatus: http.StatusForbidden,
			ExpectBody:   `blocked`,
			Sleep:        base.SleepTime,
		}),
		Entry("verify the route(should not be blocked)", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusOK,
			ExpectBody:   `hello world`,
		}),
		Entry("delete route", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/routes/r1",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("hit the route just deleted", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusNotFound,
			ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
			Sleep:        base.SleepTime,
		}),
	)

	DescribeTable("test route with plugin orchestration (post method)",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("make sure the route is not created", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusNotFound,
			ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
		}),
	)

	var routeID string
	It("create route with correct dag config by post", func() {
		resp, code, err := base.HttpPost(base.ManagerAPIHost+"/apisix/admin/routes",
			map[string]string{"Authorization": base.GetToken()}, dagConf)
		Expect(err).To(BeNil())
		Expect(code).Should(Equal(200))
		routeID = gjson.Get(string(resp), "data.id").String()
	})

	It("test the route with plugin orchestration", func() {
		base.RunTestCase(base.HttpTestCase{
			Desc:         "verify the route (should be blocked)",
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Query:        "t=root.exe",
			ExpectStatus: http.StatusForbidden,
			ExpectBody:   `blocked`,
			Sleep:        base.SleepTime,
		})
		base.RunTestCase(base.HttpTestCase{
			Desc:         "verify the route (should not be blocked)",
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusOK,
			ExpectBody:   `hello world`,
		})
		base.RunTestCase(base.HttpTestCase{
			Desc:         "delete the route by routeID",
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/routes/" + routeID,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		})
		base.RunTestCase(base.HttpTestCase{
			Desc:         "make sure the route has been deleted",
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusNotFound,
			ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
			Sleep:        base.SleepTime,
		})
	})
})
