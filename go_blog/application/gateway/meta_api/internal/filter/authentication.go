package filter

import (
	"net/http"
	"strings"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Authentication http://127.0.0.1:9000/apisix/admin/tool/version
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/apisix/admin/user/login" ||
			c.Request.URL.Path == "/apisix/admin/tool/version" ||
			!strings.HasPrefix(c.Request.URL.Path, "/apisix") {
			c.Next()
			return
		}

		cookie, _ := conf.CookieStore.Get(c.Request, "oidc")
		errResp := gin.H{
			"code":    010013,
			"message": "request unauthorized",
		}

		if cookie.IsNew {
			tokenStr := c.GetHeader("Authorization")
			// verify token
			token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(conf.AuthConf.Secret), nil
			})

			if err != nil || token == nil || !token.Valid {
				log.Warnf("token validate failed: %s", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
				return
			}

			claims, ok := token.Claims.(*jwt.StandardClaims)
			if !ok {
				log.Warnf("token validate failed: %s, %v", err, token.Valid)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
				return
			}

			if err := token.Claims.Valid(); err != nil {
				log.Warnf("token claims validate failed: %s", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
				return
			}

			if claims.Subject == "" {
				log.Warn("token claims subject empty")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
				return
			}

			if _, ok := conf.UserList[claims.Subject]; !ok {
				log.Warnf("user not exists by token claims subject %s", claims.Subject)
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
				return
			}
		} else {
			if cookie.Values["oidc_id"] != conf.OidcId {
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
				return
			}
		}

		c.Next()
	}
}
