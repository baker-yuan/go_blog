package proto_test

import (
	"encoding/json"
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var correctProtobuf = `syntax = "proto3";
    package helloworld;
    service Greeter {
        rpc SayHello (HelloRequest) returns (HelloReply) {}
    }
    message HelloRequest {
        string name = 1;
    }
    message HelloReply {
        string message = 1;
    }`

var _ = Describe("Proto", func() {
	var _ = DescribeTable("Basic",
		func(f func()) {
			f()
		},
		Entry("create proto success", func() {
			createProtoBody := make(map[string]interface{})
			createProtoBody["id"] = 1
			createProtoBody["desc"] = "test_proto1"
			createProtoBody["content"] = correctProtobuf

			_createProtoBody, err := json.Marshal(createProtoBody)
			Expect(err).To(BeNil())

			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodPost,
				Path:         "/apisix/admin/proto",
				Body:         string(_createProtoBody),
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("create proto failed, id existed", func() {
			createProtoBody := make(map[string]interface{})
			createProtoBody["id"] = 1
			createProtoBody["desc"] = "test_proto1"
			createProtoBody["content"] = correctProtobuf

			_createProtoBody, err := json.Marshal(createProtoBody)
			Expect(err).To(BeNil())

			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodPost,
				Path:         "/apisix/admin/proto",
				Body:         string(_createProtoBody),
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "proto id exists",
				ExpectStatus: http.StatusBadRequest,
			})
		}),
		Entry("update proto success", func() {
			updateProtoBody := make(map[string]interface{})
			updateProtoBody["id"] = 1
			updateProtoBody["desc"] = "test_proto1_modify"
			updateProtoBody["content"] = correctProtobuf

			_updateProtoBody, err := json.Marshal(updateProtoBody)
			Expect(err).To(BeNil())

			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodPut,
				Path:         "/apisix/admin/proto",
				Body:         string(_updateProtoBody),
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "test_proto1_modify",
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("list proto", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodGet,
				Path:         "/apisix/admin/proto",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "test_proto1_modify",
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("get proto", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodGet,
				Path:         "/apisix/admin/proto/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "test_proto1_modify",
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("delete not existed proto", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/proto/not-exist",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusNotFound,
			})
		}),
		Entry("delete proto", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/proto/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		}),
	)

	var _ = DescribeTable("Proto with grpc-transcode plugin",
		func(f func()) {
			f()
		},
		Entry("create proto success", func() {
			createProtoBody := make(map[string]interface{})
			createProtoBody["id"] = 1
			createProtoBody["desc"] = "test_proto1"
			createProtoBody["content"] = correctProtobuf

			_createProtoBody, err := json.Marshal(createProtoBody)
			Expect(err).To(BeNil())

			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodPost,
				Path:         "/apisix/admin/proto",
				Body:         string(_createProtoBody),
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("create route with grpc-transcode", func() {
			createRouteBody := make(map[string]interface{})
			createRouteBody["id"] = 1
			createRouteBody["name"] = "test_route"
			createRouteBody["uri"] = "/grpc_test"
			createRouteBody["methods"] = []string{"GET", "POST"}
			createRouteBody["upstream"] = map[string]interface{}{
				"nodes": []map[string]interface{}{
					{
						"host":   base.UpstreamGrpcIp,
						"port":   50051,
						"weight": 1,
					},
				},
				"type":   "roundrobin",
				"scheme": "grpc",
			}
			createRouteBody["plugins"] = map[string]interface{}{
				"grpc-transcode": map[string]interface{}{
					"method":   "SayHello",
					"proto_id": "1",
					"service":  "helloworld.Greeter",
				},
			}

			_createRouteBody, err := json.Marshal(createRouteBody)
			Expect(err).To(BeNil())

			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodPost,
				Path:         "/apisix/admin/routes",
				Body:         string(_createRouteBody),
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("hit GET route for grpc-transcode test", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.APISIXExpect(),
				Method:       http.MethodGet,
				Path:         "/grpc_test",
				Query:        "name=world",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "{\"message\":\"Hello world\"}",
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("hit POST route for grpc-transcode test", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.APISIXExpect(),
				Method:       http.MethodPost,
				Path:         "/grpc_test",
				Body:         "name=world",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "{\"message\":\"Hello world\"}",
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("hit JSON POST route for grpc-transcode test", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.APISIXExpect(),
				Method:       http.MethodPost,
				Path:         "/grpc_test",
				Body:         "{\"name\": \"world\"}",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "{\"message\":\"Hello world\"}",
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("delete route used proto", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/proto/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectBody:   "proto used check invalid: route",
				ExpectStatus: http.StatusBadRequest,
			})
		}),
		Entry("delete conflict route", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/routes/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		}),
		Entry("delete proto again", func() {
			base.RunTestCase(base.HttpTestCase{
				Object:       base.ManagerApiExpect(),
				Method:       http.MethodDelete,
				Path:         "/apisix/admin/proto/1",
				Headers:      map[string]string{"Authorization": base.GetToken()},
				ExpectStatus: http.StatusOK,
			})
		}),
	)
})
