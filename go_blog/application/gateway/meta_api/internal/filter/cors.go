package filter

import (
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		if conf.SecurityConf.AllowOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", conf.SecurityConf.AllowOrigin)
		}

		if conf.SecurityConf.AllowHeaders != "" {
			c.Writer.Header().Set("Access-Control-Allow-Headers", conf.SecurityConf.AllowHeaders)
		}

		if conf.SecurityConf.AllowMethods != "" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", conf.SecurityConf.AllowMethods)
		}

		if conf.SecurityConf.AllowCredentials != "" {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", conf.SecurityConf.AllowCredentials)
		}

		if conf.SecurityConf.XFrameOptions != "" {
			c.Writer.Header().Set("X-Frame-Options", conf.SecurityConf.XFrameOptions)
		}

		if conf.SecurityConf.ContentSecurityPolicy != "" {
			c.Writer.Header().Set("Content-Security-Policy", conf.SecurityConf.ContentSecurityPolicy)
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
