package retcode

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/filters/httptrace"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/codec"
	"trpc.group/trpc-go/trpc-go/log"
)

func init() {
	// 序列化
	codec.Marshaler = protojson.MarshalOptions{
		// 指定编码器是否应该以多行缩进格式输出 JSON，使得每个文本元素都在新的一行。如果 Indent 字段为空字符串，则会选择一个任意的缩进样式。
		Multiline: false,
		// 指定在多行格式化输出中使用的缩进字符集。每个条目都会以 Indent 开头并以换行符结束。如果这个字段非空，则 Multiline 被视为 true。Indent 只能由空格或制表符组成。
		Indent: "",
		// 允许包含缺失必需字段的消息进行编码而不返回错误。如果 AllowPartial 为 false（默认值），编码器在存在任何缺失必需字段时将返回错误。
		AllowPartial: true,
		// 使用 Protocol Buffers 字段的原始名称而不是 JSON 字段名的 lowerCamelCase 格式
		UseProtoNames: false,
		// 将枚举值作为数字输出
		UseEnumNumbers: true,
		// 指定是否输出未填充的字段。它不会输出未填充的 oneof 字段或扩展字段。对于未填充的字段，输出的 JSON 值如表格所示。
		EmitUnpopulated: true,
		// 指定是否输出具有默认值的基本类型字段、空列表和空映射。受影响的字段如表格所示。它的行为类似于 EmitUnpopulated，但不会输出值为 "null" 的字段，即省略的具有存在感知的字段将保持省略，以保留存在感知。EmitUnpopulated 优先于 EmitDefaultValues，因为前者生成的是后者的严格超集。
		EmitDefaultValues: false,
		// 用于在扩展 google.protobuf.Any 消息时查找类型。如果为 nil，则默认使用 protoregistry.GlobalTypes。
		// 这个结构体通常用于配置 JSON 编码器的行为，以满足特定的序列化需求。例如，你可能想要以一种易于阅读的格式输出 JSON，或者你可能需要确保所有的字段，即使它们没有被填充，也都被输出到 JSON 中。通过设置这些选项，你可以控制 JSON 编码的过程以适应你的应用程序。
		Resolver: nil,
	}
	// 反序列化
	codec.Unmarshaler = protojson.UnmarshalOptions{
		// 如果设置为 true，则在解析消息时，如果结果中缺少必需的字段，不会返回错误。这允许消息被部分填充，通常用于在不完整数据的情况下进行解析。
		AllowPartial: false,
		// 如果设置为 true，在解析过程中遇到的未知字段和枚举名称值将被忽略。这对于处理可能包含额外、未在 .proto 文件中定义的字段的 JSON 数据很有用，因为它允许解析器跳过这些字段而不报错。
		DiscardUnknown: true,
		// 用于在解析 google.protobuf.Any 消息或扩展字段时查找类型的解析器。如果为 nil，则默认使用 protoregistry.GlobalTypes。这个解析器可以帮助解析器正确处理消息中的动态类型和扩展。
		Resolver: nil,
		// 限制消息嵌套的深度。如果设置为 0，则应用默认的限制。这是一个安全特性，用于防止恶意构造的数据导致解析器递归过深，可能引起栈溢出等问题。
		RecursionLimit: 0,
	}

	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/go_blog/application/auth/cmd/trpc_go.yaml"
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

// InitServerCodec 1、自定义错误处理 2、统一返回值
func InitServerCodec(skipRspHandlerMethods []string) {
	//skipRspHandlerMethodMap := util.NewFieldExtractor[string, string]().
	//	GroupByField(skipRspHandlerMethods, func(s string) string {
	//		return s
	//	})
	//
	//// 自定义错误处理函数
	//thttp.DefaultServerCodec.ErrHandler = func(w http.ResponseWriter, r *http.Request, e *errs.Error) {
	//	ctx := r.Context()
	//	log.WarnContextf(ctx, "#errHandler catch error: %+v", e)
	//	_, _ = w.Write(NewErrorResponse(ctx, r, uint32(e.Code), e.Msg))
	//}
	//
	//// 自定义返回数据处理函数
	//thttp.DefaultServerCodec.RspHandler = func(w http.ResponseWriter, r *http.Request, rspbody []byte) error {
	//	ctx := r.Context()
	//	if len(rspbody) == 0 {
	//		log.DebugContextf(ctx, "#rspHandler skip rspbody is empty url: %s", r.RequestURI)
	//		return nil
	//	}
	//
	//	// 指定接口不进行处理
	//	if _, exist := skipRspHandlerMethodMap[r.URL.Path]; exist {
	//		log.DebugContextf(ctx, "#rspHandler hit skipRspHandlerMethods url: %s", r.RequestURI)
	//		_, err := w.Write(rspbody)
	//		return err
	//	}
	//
	//	// 检查响应是否已经是标准格式，返回的内容如果是非结构体（结构体数组、map），需要自己定义code+message+data
	//	var standardResp Response
	//	if err := json.Unmarshal(rspbody, &standardResp); err == nil &&
	//		(standardResp.Code != nil && standardResp.Message != nil && standardResp.Data != nil) {
	//		log.DebugContextf(r.Context(), "#rspHandler skip already standard response url: %s", r.RequestURI)
	//		_, err := w.Write(FillResponse(ctx, r, standardResp))
	//		return err
	//	}
	//
	//	// 统一加上code+message
	//	_, err := w.Write(NewSuccessResponse(ctx, r, rspbody))
	//	return err
	//}
}
