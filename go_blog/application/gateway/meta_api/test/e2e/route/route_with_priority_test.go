package route_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable("route with priority test",
	func(tc base.HttpTestCase) {
		base.RunTestCase(tc)
	},
	Entry("add another route with no priority (default 0)", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
					"name": "route1",
					"uri": "/server_port",
					"methods": ["GET"],
					"upstream": {
						"type": "roundrobin",
						"nodes": {
							"` + base.UpstreamIp + `:1981": 1
						}
					}
				}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("access the route", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/server_port",
		ExpectStatus: http.StatusOK,
		ExpectBody:   "1981",
		Sleep:        base.SleepTime,
	}),
	Entry("add another route with valid priority (1), upstream is different from the others", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r2",
		Body: `{
					"name": "route2",
					"uri": "/server_port",
					"methods": ["GET"],
					"priority": 1,
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
	Entry("access the route to determine whether it meets the priority (compare 1 and default)", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/server_port",
		ExpectStatus: http.StatusOK,
		ExpectBody:   "1982",
		Sleep:        base.SleepTime,
	}),
	Entry("delete route (r1)", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Method:       http.MethodDelete,
		Path:         "/apisix/admin/routes/r1",
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		Sleep:        base.SleepTime,
	}),
	Entry("delete route (r2)", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Method:       http.MethodDelete,
		Path:         "/apisix/admin/routes/r2",
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		Sleep:        base.SleepTime,
	}),
	Entry("hit the route just delete", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/server_port",
		ExpectStatus: http.StatusNotFound,
		ExpectBody:   "{\"error_msg\":\"404 Route Not Found\"}\n",
		Sleep:        base.SleepTime,
	}),
)
