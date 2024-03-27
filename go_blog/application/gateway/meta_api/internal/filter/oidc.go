package filter

import (
	"net/http"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Token struct {
	AccessToken string
}

func (token *Token) Token() (*oauth2.Token, error) {
	oauth2Token := &oauth2.Token{AccessToken: token.AccessToken}
	return oauth2Token, nil
}

func Oidc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/apisix/admin/oidc/login" {
			url := conf.OidcConfig.AuthCodeURL(conf.State)
			c.Redirect(302, url)
			c.Abort()
			return
		}

		if c.Request.URL.Path == "/apisix/admin/oidc/callback" {
			state := c.Query("state")
			if state != conf.State {
				log.Warn("the state does not match")
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			// in exchange for token
			oauth2Token, err := conf.OidcConfig.Exchange(c, c.Query("code"))
			if err != nil {
				log.Warnf("exchange code for token failed: %s", err)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			// in exchange for user's information
			token := &Token{oauth2Token.AccessToken}
			providerConfig := oidc.ProviderConfig{UserInfoURL: conf.OidcUserInfoURL}
			provider := providerConfig.NewProvider(c)
			userInfo, err := provider.UserInfo(c, token)
			if err != nil {
				log.Warnf("exchange access_token for user's information failed: %s", err)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			// set the cookie
			conf.CookieStore.MaxAge(conf.OidcExpireTime)
			cookie, _ := conf.CookieStore.Get(c.Request, "oidc")
			cookie.Values["oidc_id"] = userInfo.Subject
			conf.OidcId = userInfo.Subject
			cookie.Save(c.Request, c.Writer)
			c.AbortWithStatus(http.StatusOK)
			return
		}

		if c.Request.URL.Path == "/apisix/admin/oidc/logout" {
			cookie, _ := conf.CookieStore.Get(c.Request, "oidc")
			if cookie.IsNew {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			cookie.Options.MaxAge = -1
			cookie.Save(c.Request, c.Writer)
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
