package ret

import (
	"context"
	"encoding/json"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/proxy/biz_ctx/http"
	auth_pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
	"google.golang.org/protobuf/proto"
	"trpc.group/trpc-go/trpc-go/log"
)

func SetFilResponse(httpCtx *http.HttpContext, msg string) {
	httpCtx.Response().SetHeader("Content-Type", "application/json")
	httpCtx.Response().SetBody(NewErrorResponse(httpCtx.Context(), 500, msg))
}

func SetSucResponse(httpCtx *http.HttpContext, stdRes auth_pb.NeedStdRes, rspBody []byte) {
	httpCtx.Response().SetHeader("Content-Type", "application/json")
	var standardResp Response
	if stdRes == auth_pb.NeedStdRes_RES_REQUIRED {
		// 加上code+message
		httpCtx.Response().SetBody(NewSuccessResponse(httpCtx.Context(), rspBody))
	} else {
		httpCtx.Response().SetBody(FillResponse(httpCtx.Context(), standardResp))
	}
}

// Response 自定义返回数据
type Response struct {
	Code    *uint32         `json:"code"`            // 自定义状态码 0-成功 其他-失败
	Message *string         `json:"message"`         // 错误消息
	Trace   *string         `json:"trace,omitempty"` // 链路ID
	Time    *uint32         `json:"time"`            // 时间
	Data    json.RawMessage `json:"data"`            // 数据
}

// NewErrorResponse 创建一个带有错误信息的响应
func NewErrorResponse(ctx context.Context, code uint32, message string) []byte {
	//trace := r.Header.Get(httptrace.HeaderTraceID)
	res := Response{
		Code:    proto.Uint32(code),
		Message: proto.String(message),
		//Trace:   proto.String(trace),
		Time: proto.Uint32(uint32(time.Now().Unix())),
		Data: nil,
	}
	byteRes, _ := json.Marshal(res)
	return byteRes
}

// NewSuccessResponse 创建一个成功的响应
func NewSuccessResponse(ctx context.Context, data interface{}) []byte {
	var jsonData json.RawMessage
	var err error

	if bytes, ok := data.([]byte); ok {
		jsonData = bytes
	} else {
		jsonData, err = json.Marshal(data)
		if err != nil {
			log.ErrorContextf(ctx, "#newSuccessResponse marshal fail data: %+v, err: %+v", data, err)
			return NewErrorResponse(ctx, 500, "Internal Server Error")
		}
	}

	//trace := r.Header.Get(httptrace.HeaderTraceID)
	res := Response{
		Code:    proto.Uint32(0),
		Message: proto.String("OK"),
		//Trace:   proto.String(trace),
		Time: proto.Uint32(uint32(time.Now().Unix())),
		Data: jsonData,
	}
	byteRes, _ := json.Marshal(res)
	return byteRes
}

// FillResponse 补充公共字段
func FillResponse(ctx context.Context, standardResp Response) []byte {
	//trace := r.Header.Get(httptrace.HeaderTraceID)
	//standardResp.Trace = proto.String(trace)
	standardResp.Time = proto.Uint32(uint32(time.Now().Unix()))
	byteRes, _ := json.Marshal(standardResp)
	return byteRes
}
