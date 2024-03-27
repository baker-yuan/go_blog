package route_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable("route with plugin uri blocker",
	func(tc base.HttpTestCase) {
		base.RunTestCase(tc)
	},
	Entry("make sure the route is not created", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/*",
		ExpectStatus: http.StatusNotFound,
		ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
	}),
	Entry("create route1", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
				  "name": "route1",
				  "uri": "/*",
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
	Entry("make sure the plugin uri blocker is not worked", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/root.exe",
		ExpectStatus: http.StatusNotFound,
	}),
	Entry("update route with uri blocker", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
					  "name": "route1",
					  "uri": "/*",
					  "plugins": {
						 "uri-blocker": {
							 "block_rules": ["hello"]
						 }
					  },
					  "upstream": {
						  "type": "roundrobin",
						  "nodes": {
							  "` + base.UpstreamIp + `:1980": 1
						  }
					  }
				  }`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("verify route that block uri", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/hello",
		ExpectStatus: http.StatusForbidden,
		Sleep:        base.SleepTime,
	}),
	Entry("update route with uri blocker", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
					  "name": "route1",
					  "uri": "/*",
					  "plugins": {
						 "uri-blocker": {
							 "block_rules": ["robots.txt"]
						 }
					  },
					  "upstream": {
						  "type": "roundrobin",
						  "nodes": {
							  "` + base.UpstreamIp + `:1980": 1
						  }
					  }
				  }`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("verify route that old block uri rule", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/hello",
		ExpectStatus: http.StatusOK,
		Sleep:        base.SleepTime,
	}),
	Entry("verify route that block uri", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/robots.txt",
		ExpectStatus: http.StatusForbidden,
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
