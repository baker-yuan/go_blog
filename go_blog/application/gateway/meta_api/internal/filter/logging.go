package filter

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start, host, remoteIP, path, method := time.Now(), c.Request.Host, c.ClientIP(), c.Request.URL.Path, c.Request.Method
		query := c.Request.URL.RawQuery
		requestId := c.Writer.Header().Get("X-Request-Id")

		blw := &bodyLogWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		latency := time.Since(start) / 1000000
		statusCode := c.Writer.Status()
		// respBody := blw.body.String()

		var errs []string
		for _, err := range c.Errors {
			errs = append(errs, err.Error())
		}

		logger.Desugar().Info(path,
			// zap.String("path", path),
			zap.Int("status", statusCode),
			zap.String("host", host),
			zap.String("query", query),
			zap.String("requestId", requestId),
			zap.Duration("latency", latency),
			zap.String("remoteIP", remoteIP),
			zap.String("method", method),
			// zap.String("respBody", respBody),
			zap.Strings("errs", errs),
		)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
