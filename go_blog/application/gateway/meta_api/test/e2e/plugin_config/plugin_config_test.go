package plugin_config_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable("Plugin Config",
	func(tc base.HttpTestCase) {
		base.RunTestCase(tc)
	},
	Entry("make sure the route doesn't exist", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/hello",
		ExpectStatus: http.StatusNotFound,
		ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
	}),
	Entry("create plugin config", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Path:   "/apisix/admin/plugin_configs/1",
		Method: http.MethodPut,
		Body: `{
				"plugins": {
					"response-rewrite": {
						"headers": {
							"X-VERSION":"1.0"
						}
					},
					"uri-blocker": {
						"block_rules": ["select.+(from|limit)", "(?:(union(.*?)select))"]
					}
				},
				"labels": {
					"version": "v1",
					"build":   "16"
				}
			}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("create plugin config by Post", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Path:   "/apisix/admin/plugin_configs",
		Method: http.MethodPost,
		Body: `{
				"id": "2",
				"plugins": {
					"response-rewrite": {
						"headers": {
							"X-VERSION":"22.0"
						}
					}
				},
				"labels": {
					"version": "v2",
					"build":   "17",
					"extra":   "test"
				}
			}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("get plugin config", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Path:         "/apisix/admin/plugin_configs/1",
		Method:       http.MethodGet,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		ExpectBody:   `"plugins":{"response-rewrite":{"headers":{"X-VERSION":"1.0"}},"uri-blocker":{"block_rules":["select.+(from|limit)","(?:(union(.*?)select))"]}}`,
		Sleep:        base.SleepTime,
	}),
	Entry("search plugin_config list by label ", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Path:         "/apisix/admin/plugin_configs",
		Query:        "label=build:16",
		Method:       http.MethodGet,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		ExpectBody:   `"labels":{"build":"16","version":"v1"}`,
		Sleep:        base.SleepTime,
	}),
	Entry("search plugin_config list by label (only key)", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Path:         "/apisix/admin/plugin_configs",
		Query:        "label=extra",
		Method:       http.MethodGet,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		ExpectBody:   `"labels":{"build":"17","extra":"test","version":"v2"}`,
		Sleep:        base.SleepTime,
	}),
	Entry("create route with the plugin config created before", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Method: http.MethodPut,
		Path:   "/apisix/admin/routes/r1",
		Body: `{
				 "name": "route1",
				 "uri": "/hello",
				 "plugin_config_id": "1",
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
	Entry("verify route with header", base.HttpTestCase{
		Object:        base.APISIXExpect(),
		Method:        http.MethodGet,
		Path:          "/hello",
		ExpectStatus:  http.StatusOK,
		ExpectBody:    "hello world",
		ExpectHeaders: map[string]string{"X-VERSION": "1.0"},
		Sleep:         base.SleepTime,
	}),
	Entry("verify route that should be blocked", base.HttpTestCase{
		Object:        base.APISIXExpect(),
		Method:        http.MethodGet,
		Path:          "/hello",
		Query:         "name=%3Bselect%20from%20sys",
		ExpectStatus:  http.StatusForbidden,
		ExpectHeaders: map[string]string{"X-VERSION": "1.0"},
		Sleep:         base.SleepTime,
	}),
	Entry("update plugin config by patch", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Path:   "/apisix/admin/plugin_configs/1",
		Method: http.MethodPatch,
		Body: `{
				"plugins": {
					"response-rewrite": {
						"headers": {
							"X-VERSION":"2.0"
						}
					},
					"uri-blocker": {
						"block_rules": ["none"]
					}
				}
			}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("verify patch update", base.HttpTestCase{
		Object:        base.APISIXExpect(),
		Method:        http.MethodGet,
		Path:          "/hello",
		ExpectStatus:  http.StatusOK,
		ExpectBody:    "hello world",
		ExpectHeaders: map[string]string{"X-VERSION": "2.0"},
		Sleep:         base.SleepTime,
	}),
	Entry("verify patch update(should not block)", base.HttpTestCase{
		Object:        base.APISIXExpect(),
		Method:        http.MethodGet,
		Path:          "/hello",
		Query:         "name=%3Bselect%20from%20sys",
		ExpectStatus:  http.StatusOK,
		ExpectBody:    "hello world",
		ExpectHeaders: map[string]string{"X-VERSION": "2.0"},
	}),
	Entry("update plugin config by sub path patch", base.HttpTestCase{
		Object: base.ManagerApiExpect(),
		Path:   "/apisix/admin/plugin_configs/1/plugins",
		Method: http.MethodPatch,
		Body: `{
				"response-rewrite": {
					"headers": {
						"X-VERSION":"3.0"
					}
				}
			}`,
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("verify patch (sub path)", base.HttpTestCase{
		Object:        base.APISIXExpect(),
		Method:        http.MethodGet,
		Path:          "/hello",
		ExpectStatus:  http.StatusOK,
		ExpectBody:    "hello world",
		ExpectHeaders: map[string]string{"X-VERSION": "3.0"},
		Sleep:         base.SleepTime,
	}),
	Entry("delete route", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Method:       http.MethodDelete,
		Path:         "/apisix/admin/routes/r1",
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
	}),
	Entry("delete plugin config", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Method:       http.MethodDelete,
		Path:         "/apisix/admin/plugin_configs/1",
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusOK,
		Sleep:        base.SleepTime,
	}),
	Entry("make sure the plugin config has been deleted", base.HttpTestCase{
		Object:       base.ManagerApiExpect(),
		Method:       http.MethodGet,
		Path:         "/apisix/admin/plugin_configs/1",
		Headers:      map[string]string{"Authorization": base.GetToken()},
		ExpectStatus: http.StatusNotFound,
		ExpectBody:   `{"code":10001,"message":"data not found"`,
		Sleep:        base.SleepTime,
	}),
	Entry("make sure the route has been deleted", base.HttpTestCase{
		Object:       base.APISIXExpect(),
		Method:       http.MethodGet,
		Path:         "/hello",
		ExpectStatus: http.StatusNotFound,
		ExpectBody:   `{"error_msg":"404 Route Not Found"}`,
		Sleep:        base.SleepTime,
	}),
)
