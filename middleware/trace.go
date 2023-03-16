package middleware

// import (
// 	"bytes"
// 	"io"
// 	"time"
//
// 	customLog "github.com/baker-yuan/go-blog/common/log"
// 	ginzap "github.com/gin-contrib/zap"
// 	"github.com/gin-gonic/gin"
// 	"go.opentelemetry.io/otel/trace"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )
//
// // TraceMiddleware 出入参、链路日志
// func TraceMiddleware() gin.HandlerFunc {
// 	return ginzap.GinzapWithConfig(customLog.GetZipLogger(), &ginzap.Config{
// 		UTC:        true,
// 		TimeFormat: time.RFC3339,
// 		Context: func(ginCtx *gin.Context) []zapcore.Field {
// 			var (
// 				fields    []zapcore.Field
// 				requestID string
// 				tee       io.Reader
// 				body      []byte
// 				buf       bytes.Buffer
// 			)
// 			// log request ID
// 			if requestID = ginCtx.Writer.Header().Get("X-Request-Id"); requestID != "" {
// 				fields = append(fields, zap.String("request_id", requestID))
// 			}
// 			// log trace and span ID
// 			if trace.SpanFromContext(ginCtx.Request.Context()).SpanContext().IsValid() {
// 				fields = append(fields, zap.String("trace_id", trace.SpanFromContext(ginCtx.Request.Context()).SpanContext().TraceID().String()))
// 				fields = append(fields, zap.String("span_id", trace.SpanFromContext(ginCtx.Request.Context()).SpanContext().SpanID().String()))
// 			}
// 			// log request body
// 			tee = io.TeeReader(ginCtx.Request.Body, &buf)
// 			body, _ = io.ReadAll(tee)
// 			ginCtx.Request.Body = io.NopCloser(&buf)
// 			fields = append(fields, zap.String("body", string(body)))
// 			return fields
// 		},
// 	})
// }
