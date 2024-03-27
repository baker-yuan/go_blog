package route_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("route with plugin cors", func() {
	DescribeTable("test route with plugin cors",
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
		Entry("create route with cors default setting", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
				"name": "route1",
				"uri": "/hello",
				"plugins": {
					"cors": {}
				},
				"upstream": {
					"type": "roundrobin",
					"nodes": [{
						"host": "` + base.UpstreamIp + `",
						"port": 1981,
						"weight": 1
					}]
				}
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route with cors default setting", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusOK,
			ExpectHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "*",
			},
			ExpectBody: "hello world",
			Sleep:      base.SleepTime,
		}),
		Entry("update route with specified setting", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
				"name": "route1",
				"uri": "/hello",
					"plugins": {
						"cors": {
							"allow_origins": "http://sub.domain.com,http://sub2.domain.com",
							"allow_methods": "GET,POST",
							"allow_headers": "headr1,headr2",
							"expose_headers": "ex-headr1,ex-headr2",
							"max_age": 50,
							"allow_credential": true
						}
					},
					"upstream": {
						"type": "roundrobin",
						"nodes": [{
							"host": "` + base.UpstreamIp + `",
							"port": 1981,
							"weight": 1
						}]
					}
				}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route with cors specified setting", base.HttpTestCase{
			Object: base.APISIXExpect(),
			Method: http.MethodGet,
			Path:   "/hello",
			Headers: map[string]string{
				"Origin":    "http://sub2.domain.com",
				"resp-vary": "Via",
			},
			ExpectStatus: http.StatusOK,
			ExpectHeaders: map[string]string{
				"Access-Control-Allow-Origin":   "http://sub2.domain.com",
				"Access-Control-Allow-Methods":  "GET,POST",
				"Access-Control-Allow-Headers":  "headr1,headr2",
				"Access-Control-Expose-Headers": "ex-headr1,ex-headr2",
				"Access-Control-Max-Age":        "50",
			},
			ExpectBody: "hello world",
			Sleep:      base.SleepTime,
		}),
		Entry("verify route with cors specified no match origin", base.HttpTestCase{
			Object: base.APISIXExpect(),
			Method: http.MethodGet,
			Path:   "/hello",
			Headers: map[string]string{
				"Origin": "http://sub3.domain.com",
			},
			ExpectStatus: http.StatusOK,
			ExpectHeaders: map[string]string{
				"Access-Control-Allow-Origin":   "",
				"Access-Control-Allow-Methods":  "",
				"Access-Control-Allow-Headers":  "",
				"Access-Control-Expose-Headers": "",
				"Access-Control-Max-Age":        "",
			},
			ExpectBody: "hello world",
			Sleep:      base.SleepTime,
		}),
		Entry("verify route with options method", base.HttpTestCase{
			Object: base.APISIXExpect(),
			Method: http.MethodOptions,
			Headers: map[string]string{
				"Origin": "http://sub2.domain.com",
			},
			Path:         "/hello",
			ExpectStatus: http.StatusOK,
			ExpectHeaders: map[string]string{
				"Access-Control-Allow-Origin":   "http://sub2.domain.com",
				"Access-Control-Allow-Methods":  "GET,POST",
				"Access-Control-Allow-Headers":  "headr1,headr2",
				"Access-Control-Expose-Headers": "ex-headr1,ex-headr2",
				"Access-Control-Max-Age":        "50",
			},
			ExpectBody: "",
		}),
		Entry("update route with cors setting force wildcard", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
				"name": "route1",
				"uri": "/hello",
				"plugins": {
					"cors": {
						"allow_origins": "**",
						"allow_methods": "**",
						"allow_headers": "**",
						"expose_headers": "*"
					}
				},
				"upstream": {
					"type": "roundrobin",
					"nodes": [{
						"host": "` + base.UpstreamIp + `",
						"port": 1981,
						"weight": 1
					}]
				}
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route with cors setting force wildcard", base.HttpTestCase{
			Object: base.APISIXExpect(),
			Method: http.MethodGet,
			Path:   "/hello",
			Headers: map[string]string{
				"Origin":                         "https://sub.domain.com",
				"ExternalHeader1":                "val",
				"ExternalHeader2":                "val",
				"ExternalHeader3":                "val",
				"Access-Control-Request-Headers": "req-header1,req-header2",
			},
			ExpectStatus: http.StatusOK,
			ExpectHeaders: map[string]string{
				"Access-Control-Allow-Origin":      "https://sub.domain.com",
				"Vary":                             "Origin",
				"Access-Control-Allow-Methods":     "GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS,CONNECT,TRACE",
				"Access-Control-Allow-Headers":     "req-header1,req-header2",
				"Access-Control-Expose-Headers":    "*",
				"Access-Control-Allow-Credentials": "",
			},
			ExpectBody: "hello world",
			Sleep:      base.SleepTime,
		}),
		Entry("delete route", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/routes/r1",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("make sure the route deleted", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			ExpectStatus: http.StatusNotFound,
			ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
			Sleep:        base.SleepTime,
		}),
	)
})
