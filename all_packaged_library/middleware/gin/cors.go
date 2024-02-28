package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetCorsRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "lang", "token", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}

// func Cors() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		method := context.Request.Method
//		context.Header("Access-Control-Allow-Origin", "*")
//		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
//		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
//		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		context.Header("Access-Control-Allow-Credentials", "true")
//		// 允许放行OPTIONS请求
//		if method == "OPTIONS" {
//			context.AbortWithStatus(http.StatusNoContent)
//		}
//		context.Next()
//	}
// }
