package filter

import (
	"net/http"
	"testing"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func genToken(username string, issueAt, expireAt int64) string {
	claims := jwt.StandardClaims{
		Subject:   username,
		IssuedAt:  issueAt,
		ExpiresAt: expireAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(conf.AuthConf.Secret))

	return signedToken
}

func TestAuthenticationMiddleware_Handle(t *testing.T) {
	r := gin.New()
	r.Use(Authentication())
	r.GET("/*path", func(c *gin.Context) {
	})

	w := performRequest(r, "GET", "/apisix/admin/user/login", nil)
	assert.Equal(t, http.StatusOK, w.Code)

	w = performRequest(r, "GET", "/apisix/admin/routes", nil)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// test with token expire
	expireToken := genToken("admin", time.Now().Unix(), time.Now().Unix()-60*3600)
	w = performRequest(r, "GET", "/apisix/admin/routes", map[string]string{"Authorization": expireToken})
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// test with empty subject
	emptySubjectToken := genToken("", time.Now().Unix(), time.Now().Unix()+60*3600)
	w = performRequest(r, "GET", "/apisix/admin/routes", map[string]string{"Authorization": emptySubjectToken})
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// test token with nonexistent username
	nonexistentUserToken := genToken("user1", time.Now().Unix(), time.Now().Unix()+60*3600)
	w = performRequest(r, "GET", "/apisix/admin/routes", map[string]string{"Authorization": nonexistentUserToken})
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// test auth success
	validToken := genToken("admin", time.Now().Unix(), time.Now().Unix()+60*3600)
	w = performRequest(r, "GET", "/apisix/admin/routes", map[string]string{"Authorization": validToken})
	assert.Equal(t, http.StatusOK, w.Code)
}
