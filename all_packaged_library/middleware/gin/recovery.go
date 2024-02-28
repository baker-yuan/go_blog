package middleware

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/gin-gonic/gin"
)

// SetRecovery 捕获所有panic，并且返回错误信息
func SetRecovery(r *gin.Engine) {
	handlerFunc := func(ginContext *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error(context.Background(), "fail err: %+v", err)
				ResponseError(ginContext)
				// 先做一下日志记录
				// fmt.Println(string(debug.Stack()))
				// public.ComLogNotice(ginContext, "_com_panic", map[string]interface{}{
				// 	"error": fmt.Sprint(err),
				// 	"stack": string(debug.Stack()),
				// })
				// if lib.ConfBase.DebugMode != "debug" {
				// 	ResponseError(ginContext, 500, errors.New("内部错误"))
				// 	return
				// } else {
				// 	ResponseError(ginContext, 500, errors.New(fmt.Sprint(err)))
				// 	return
				// }
			}
		}()
		ginContext.Next()
	}
	r.Use(handlerFunc)
}
