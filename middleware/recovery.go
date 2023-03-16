package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
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
}
