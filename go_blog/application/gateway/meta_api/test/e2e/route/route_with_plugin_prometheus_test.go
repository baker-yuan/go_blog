package route_test

import (
	"net/http"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable("route with plugin prometheus",
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
	Entry("create route with plugin prometheus", base.HttpTestCase{
		Desc:   "create route",
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
				"name": "route1",
				"uri": "/hello",
				"plugins": {
					"prometheus": {}
				},
				"upstream": {
					"type": "roundrobin",
					"nodes": {
						"` + base.UpstreamIp + `:1982": 1
					}
				}
			}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("fetch the prometheus metric data", base.HttpTestCase{
		Object:       base.PrometheusExporterExpect(),
		Method:       http.MethodGet,
		Path:         "/apisix/prometheus/metrics",
		ExpectStatus: http.StatusOK,
		ExpectBody:   "apisix_etcd_reachable 1",
		Sleep:        base.SleepTime,
	}),
	Entry("request from client (200)", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/hello",
		ExpectStatus: http.StatusOK,
		ExpectBody:   "hello world",
	}),
	Entry("create route that uri not exists in upstream", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
				"name": "route1",
				"uri": "/hello-not-exists",
				"plugins": {
					"prometheus": {}
				},
				"upstream": {
					"type": "roundrobin",
					"nodes": {
						"` + base.UpstreamIp + `:1982": 1
					}
				}
			}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("request from client (404)", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/hello-not-exists",
		ExpectStatus: http.StatusNotFound,
		Sleep:        base.SleepTime,
	}),
	Entry("verify the prometheus metric data (apisix_http_status 200)", base.HttpTestCase{
		Object:       base.PrometheusExporterExpect(),
		Method:       http.MethodGet,
		Path:         "/apisix/prometheus/metrics",
		ExpectStatus: http.StatusOK,
		ExpectBody:   `apisix_http_status{code="200",route="r1",matched_uri="/hello",matched_host="",service="",consumer=""`,
		Sleep:        1 * time.Second,
	}),
	Entry("verify the prometheus metric data (apisix_http_status 404)", base.HttpTestCase{
		Object:       base.PrometheusExporterExpect(),
		Method:       http.MethodGet,
		Path:         "/apisix/prometheus/metrics",
		ExpectStatus: http.StatusOK,
		ExpectBody:   `apisix_http_status{code="404",route="r1",matched_uri="/hello-not-exists",matched_host="",service="",consumer=""`,
		Sleep:        base.SleepTime,
	}),
	Entry("delete route", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Method:       http.MethodDelete,
		Path:         "/apisix/admin/routes/r1",
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		Sleep:        base.SleepTime,
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
