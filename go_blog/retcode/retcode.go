package retcode

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/common/util"
	"github.com/baker-yuan/go-blog/go_blog/filters/httptrace"
	"google.golang.org/protobuf/proto"
	"trpc.group/trpc-go/trpc-go/errs"
	thttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

// Response 自定义返回数据
type Response struct {
	Code    *uint32         `json:"code"`            // 自定义状态码 0-成功 其他-失败
	Message *string         `json:"message"`         // 错误消息
	Trace   *string         `json:"trace,omitempty"` // 链路ID
	Time    *uint32         `json:"time"`            // 时间
	Data    json.RawMessage `json:"data"`            // 数据
}

// NewErrorResponse 创建一个带有错误信息的响应
func NewErrorResponse(ctx context.Context, r *http.Request, code uint32, message string) []byte {
	trace := r.Header.Get(httptrace.HeaderTraceID)
	res := Response{
		Code:    proto.Uint32(code),
		Message: proto.String(message),
		Trace:   proto.String(trace),
		Time:    proto.Uint32(uint32(time.Now().Unix())),
		Data:    nil,
	}
	byteRes, _ := json.Marshal(res)
	return byteRes
}

// NewSuccessResponse 创建一个成功的响应
func NewSuccessResponse(ctx context.Context, r *http.Request, data interface{}) []byte {
	var jsonData json.RawMessage
	var err error

	if bytes, ok := data.([]byte); ok {
		jsonData = bytes
	} else {
		jsonData, err = json.Marshal(data)
		if err != nil {
			log.ErrorContextf(ctx, "#newSuccessResponse marshal fail data: %+v, err: %+v", data, err)
			return NewErrorResponse(ctx, r, 500, "Internal Server Error")
		}
	}

	trace := r.Header.Get(httptrace.HeaderTraceID)
	res := Response{
		Code:    proto.Uint32(0),
		Message: proto.String("OK"),
		Trace:   proto.String(trace),
		Time:    proto.Uint32(uint32(time.Now().Unix())),
		Data:    jsonData,
	}
	byteRes, _ := json.Marshal(res)
	return byteRes
}

// FillResponse 补充公共字段
func FillResponse(ctx context.Context, r *http.Request, standardResp Response) []byte {
	trace := r.Header.Get(httptrace.HeaderTraceID)
	standardResp.Trace = proto.String(trace)
	standardResp.Time = proto.Uint32(uint32(time.Now().Unix()))
	byteRes, _ := json.Marshal(standardResp)
	return byteRes
}

// InitServerCodec
// https://iwiki.woa.com/p/490796254
// https://mk.woa.com/q/281362
func InitServerCodec(skipRspHandlerMethods []string) {

	skipRspHandlerMethodMap := util.NewFieldExtractor[string, string]().
		GroupByField(skipRspHandlerMethods, func(s string) string {
			return s
		})

	// 自定义错误处理函数
	thttp.DefaultServerCodec.ErrHandler = func(w http.ResponseWriter, r *http.Request, e *errs.Error) {
		ctx := r.Context()
		log.WarnContextf(ctx, "#errHandler catch error: %+v", e)
		_, _ = w.Write(NewErrorResponse(ctx, r, uint32(e.Code), e.Msg))
	}

	// 自定义返回数据处理函数
	thttp.DefaultServerCodec.RspHandler = func(w http.ResponseWriter, r *http.Request, rspbody []byte) error {
		ctx := r.Context()
		if len(rspbody) == 0 {
			log.DebugContextf(ctx, "#rspHandler skip rspbody is empty url: %s", r.RequestURI)
			return nil
		}

		// 指定接口不进行处理
		if _, exist := skipRspHandlerMethodMap[r.URL.Path]; exist {
			log.DebugContextf(ctx, "#rspHandler hit skipRspHandlerMethods url: %s", r.RequestURI)
			_, err := w.Write(rspbody)
			return err
		}

		// 检查响应是否已经是标准格式，返回的内容如果是非结构体（结构体数组、map），需要自己定义code+message+data
		var standardResp Response
		if err := json.Unmarshal(rspbody, &standardResp); err == nil &&
			(standardResp.Code != nil && standardResp.Message != nil && standardResp.Data != nil) {
			log.DebugContextf(r.Context(), "#rspHandler skip already standard response url: %s", r.RequestURI)
			_, err := w.Write(FillResponse(ctx, r, standardResp))
			return err
		}

		// 统一加上code+message
		_, err := w.Write(NewSuccessResponse(ctx, r, rspbody))
		return err
	}
}
