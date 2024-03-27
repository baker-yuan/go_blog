package system_config_test

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("system config", func() {
	DescribeTable("test system config data CURD",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},

		Entry("get system config should get not found error", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/system_config/grafana",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
		}),

		Entry("create system config should get schema validate failed error", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPost,
			Path:   "/apisix/admin/system_config",
			Body: `{
				"config_name": "",
				"payload": {"url":"http://127.0.0.1:3000"}
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusBadRequest,
		}),

		Entry("create system config should success", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPost,
			Path:   "/apisix/admin/system_config",
			Body: `{
				"config_name": "grafana",
				"payload": {"url":"http://127.0.0.1:3000"}
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "\"config_name\":\"grafana\",\"payload\":{\"url\":\"http://127.0.0.1:3000\"}",
		}),

		Entry("after create system config get config should succeed", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/system_config/grafana",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "\"config_name\":\"grafana\",\"payload\":{\"url\":\"http://127.0.0.1:3000\"}",
		}),

		Entry("update system config should get schema validate failed error", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/system_config",
			Body: `{
				"config_name": "",
				"payload": {"url":"http://127.0.0.1:2000"}
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusBadRequest,
		}),

		Entry("update system config should success", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/system_config",
			Body: `{
				"config_name": "grafana",
				"payload": {"url":"http://127.0.0.1:2000"}
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "\"config_name\":\"grafana\",\"payload\":{\"url\":\"http://127.0.0.1:2000\"}",
		}),

		Entry("after update system config get config should succeed", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/system_config/grafana",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "\"config_name\":\"grafana\",\"payload\":{\"url\":\"http://127.0.0.1:2000\"}",
		}),

		Entry("delete system config should success", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/system_config/grafana",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),

		Entry("get system config should get not found error", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodGet,
			Path:         "/apisix/admin/system_config/grafana",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
		}),
	)
})
