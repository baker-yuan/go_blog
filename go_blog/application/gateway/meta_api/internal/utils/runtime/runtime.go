package runtime

import (
	"net/http"
	"runtime"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
)

var (
	ActuallyPanic = true
)

var PanicHandlers = []func(interface{}){logPanic}

func HandlePanic(additionalHandlers ...func(interface{})) {
	if err := recover(); err != nil {
		for _, fn := range PanicHandlers {
			fn(err)
		}
		for _, fn := range additionalHandlers {
			fn(err)
		}
		if ActuallyPanic {
			panic(err)
		}
	}
}

func logPanic(r interface{}) {
	if r == http.ErrAbortHandler {
		return
	}

	const size = 32 << 10
	stacktrace := make([]byte, size)
	stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]
	if _, ok := r.(string); ok {
		log.Errorf("observed a panic: %s\n%s", r, stacktrace)
	} else {
		log.Errorf("observed a panic: %#v (%v)\n%s", r, r, stacktrace)
	}
}
