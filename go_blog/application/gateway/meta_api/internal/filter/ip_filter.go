package filter

import (
	"net"
	"net/http"
	"strings"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils/consts"
	"github.com/gin-gonic/gin"
)

type subnet struct {
	ipStr   string
	ipNet   *net.IPNet
	allowed bool
}

func generateIPSet(ipList []string) (map[string]bool, []*subnet) {
	var ips = map[string]bool{}
	var subnets []*subnet
	for _, ipStr := range ipList {
		if ip, net, err := net.ParseCIDR(ipStr); err == nil {
			if n, total := net.Mask.Size(); n == total {
				ips[ip.String()] = true
				continue
			}

			subnets = append(subnets, &subnet{
				ipStr:   ipStr,
				ipNet:   net,
				allowed: true,
			})
			continue
		}
		if ip := net.ParseIP(ipStr); ip != nil {
			ips[ip.String()] = true
		}
	}

	return ips, subnets
}

func checkIP(ipStr string, ips map[string]bool, subnets []*subnet) bool {
	allowed, ok := ips[ipStr]
	if ok {
		return allowed
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	for _, subnet := range subnets {
		if subnet.ipNet.Contains(ip) {
			return subnet.allowed
		}
	}

	return false
}

func IPFilter() gin.HandlerFunc {
	ips, subnets := generateIPSet(conf.AllowList)
	return func(c *gin.Context) {
		var ipStr string
		if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
			ipStr = ip
		}

		if len(conf.AllowList) < 1 {
			c.Next()
			return
		}

		if ipStr == "" {
			log.Warn("forbidden by empty IP")
			c.AbortWithStatusJSON(http.StatusForbidden, consts.ErrIPNotAllow)
			return
		}

		res := checkIP(ipStr, ips, subnets)
		if !res {
			log.Warnf("forbidden by IP: %s, allowed list: %v", ipStr, conf.AllowList)
			c.AbortWithStatusJSON(http.StatusForbidden, consts.ErrIPNotAllow)
		}

		c.Next()
	}
}
