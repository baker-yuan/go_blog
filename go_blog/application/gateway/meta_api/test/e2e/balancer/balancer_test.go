package balancer_test

import (
	"net/http"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Balancer", func() {
	DescribeTable("test create upstream and route",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("create upstream (roundrobin with same weight)", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/upstreams/1",
			Body: `{
				"nodes": [{
					"host": "` + base.UpstreamIp + `",
					"port": 1980,
					"weight": 1
				},
				{
					"host": "` + base.UpstreamIp + `",
					"port": 1981,
					"weight": 1
				},
				{
					"host": "` + base.UpstreamIp + `",
					"port": 1982,
					"weight": 1
				}],
				"type": "roundrobin"
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("create route using the upstream just created", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/1",
			Body: `{
				"name": "route1",
				"uri": "/server_port",
				"upstream_id": "1"
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			Sleep:        base.SleepTime,
		}),
	)
	It("verify balancer by access count(same weight)", func() {
		time.Sleep(base.SleepTime)
		// batch test /server_port api
		res := base.BatchTestServerPort(18, nil, "")
		assert.Equal(GinkgoT(), 6, res["1980"])
		assert.Equal(GinkgoT(), 6, res["1981"])
		assert.Equal(GinkgoT(), 6, res["1982"])
	})

	DescribeTable("test update upstream",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("update upstream (roundrobin with different weight)", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/upstreams/1",
			Body: `{
				"nodes": [{
					"host": "` + base.UpstreamIp + `",
					"port": 1980,
					"weight": 1
				},
				{
					"host": "` + base.UpstreamIp + `",
					"port": 1981,
					"weight": 2
				},
				{
					"host": "` + base.UpstreamIp + `",
					"port": 1982,
					"weight": 3
				}],
				"type": "roundrobin"
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
	)
	It("verify balancer by access count(different weight)", func() {
		time.Sleep(base.SleepTime)
		// batch test /server_port api
		res := base.BatchTestServerPort(18, nil, "")
		assert.Equal(GinkgoT(), 3, res["1980"])
		assert.Equal(GinkgoT(), 6, res["1981"])
		assert.Equal(GinkgoT(), 9, res["1982"])
	})

	DescribeTable("update upstream",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("update upstream (roundrobin with weight 1 and 0)", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/upstreams/1",
			Body: `{
				"nodes": [{
					"host": "` + base.UpstreamIp + `",
					"port": 1980,
					"weight": 1
				},
				{
					"host": "` + base.UpstreamIp + `",
					"port": 1981,
					"weight": 0
				}],
				"type": "roundrobin"
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
	)
	It("verify balancer by access count(weight 1 and 0)", func() {
		time.Sleep(base.SleepTime)
		// batch test /server_port api
		res := base.BatchTestServerPort(18, nil, "")
		assert.Equal(GinkgoT(), 18, res["1980"])
	})

	DescribeTable("update upstream",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("update upstream (roundrobin with weight only 1)", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/upstreams/1",
			Body: `{
				"nodes": [{
					"host": "` + base.UpstreamIp + `",
					"port": 1980,
					"weight": 1
				}],
				"type": "roundrobin"
			}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
	)
	It("verify balancer by access count(weight only 1)", func() {
		time.Sleep(base.SleepTime)
		// batch test /server_port api
		res := base.BatchTestServerPort(18, nil, "")
		assert.Equal(GinkgoT(), 18, res["1980"])
	})

	Context("test balancer delete", func() {
		It("delete route", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/routes/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		})
		It("delete upstream", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/upstreams/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		})
		It("hit the route just deleted", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.APISIXExpect(),
				Method:       http.MethodGet,
				Path:         "/server_port",
				ExpectStatus: http.StatusNotFound,
				ExpectBody:   "{\"error_msg\":\"404 Route Not Found\"}\n",
				Sleep:        base.SleepTime,
			})
		})
	})
})
