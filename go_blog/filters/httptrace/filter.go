package httptrace

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

func init() {
	filter.Register("httpTrace", ServerFilter(), nil)
}

const HeaderTraceID = "header-trace-id"

// ServerFilter 尝试从context取处traceID并写入header
func ServerFilter() filter.ServerFilter {
	return func(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (interface{}, error) {
		rsp, err := next(ctx, req)

		spanContext := trace.SpanContextFromContext(ctx)
		log.DebugContextf(ctx, "requestFilter : traceID is %v", spanContext.TraceID().String())

		header := http.Head(ctx)
		if header != nil {
			header.Request.Header.Add(HeaderTraceID, spanContext.TraceID().String())
		}

		return rsp, err
	}
}
