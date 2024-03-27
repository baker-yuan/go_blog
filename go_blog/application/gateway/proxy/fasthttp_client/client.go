package fasthttp_client

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/valyala/fasthttp"
)

// ProxyTimeout 发送HTTP请求，并支持超时控制
// @scheme 请求http协议
// @node 下游服务节点信息
func ProxyTimeout(scheme string, node biz_ctx.IInstance, req *fasthttp.Request, resp *fasthttp.Response, timeout time.Duration) error {
	// 官网案例，得到的就是：http://demo.apinto.com:8280
	addr := fmt.Sprintf("%s://%s", scheme, node.Addr())
	err := defaultClient.ProxyTimeout(addr, req, resp, timeout)
	if err != nil {
		node.Down()
	}
	return err
}

var (
	defaultClient Client
)

const (
	DefaultMaxConns           = 10240
	DefaultMaxConnWaitTimeout = time.Second * 60
)

// Client implements http client.
//
// Copying Client by value is prohibited. Create new instance instead.
//
// It is safe calling Client methods from concurrently running goroutines.
//
// The fields of a Client should not be changed while it is in use.
type Client struct {
	mLock sync.Mutex                      // 初始化 m ms 加锁
	m     map[string]*fasthttp.HostClient // HTTP的HostClient key=host(ip:port) value=fasthttp.HostClient
	ms    map[string]*fasthttp.HostClient // HTTPS的HostClient key=host(ip:port) value=fasthttp.HostClient
}

func (c *Client) getHostClient(addr string) (*fasthttp.HostClient, string, error) {
	// 从地址中获取http协议
	scheme, host := readAddress(addr)

	// 是否是https协议
	isTLS := false
	if strings.EqualFold(scheme, "https") {
		isTLS = true
	} else if !strings.EqualFold(scheme, "http") {
		return nil, "", fmt.Errorf("unsupported protocol %q. http and https are supported", scheme)
	}

	// 是否需要清理空闲连接
	startCleaner := false

	c.mLock.Lock()

	m := c.m
	if isTLS {
		m = c.ms
	}

	// 初始化map
	if m == nil {
		m = make(map[string]*fasthttp.HostClient)
		if isTLS {
			c.ms = m
		} else {
			c.m = m
		}
	}

	hc := m[host]

	// 初始化fasthttp.HostClient
	if hc == nil {
		hc = &fasthttp.HostClient{
			Addr:               addMissingPort(host, isTLS),
			IsTLS:              isTLS,
			Dial:               Dial,
			MaxConns:           DefaultMaxConns,
			MaxConnWaitTimeout: DefaultMaxConnWaitTimeout,
			RetryIf: func(request *fasthttp.Request) bool {
				return false
			},
		}
		m[host] = hc
		if len(m) == 1 {
			startCleaner = true
		}
	}
	c.mLock.Unlock()

	if startCleaner {
		go c.mCleaner(m)
	}
	return hc, scheme, nil
}

// ProxyTimeout performs the given request and waits for response during
// the given timeout duration.
//
// Request must contain at least non-zero RequestURI with full url (including
// scheme and host) or non-zero Host header + RequestURI.
//
// Client determines the server to be requested in the following order:
//
//   - from RequestURI if it contains full url with scheme and host;
//   - from Host header otherwise.
//
// The function doesn't follow redirects. Use Get* for following redirects.
//
// Response is ignored if resp is nil.
//
// ErrTimeout is returned if the response wasn't returned during
// the given timeout.
//
// ErrNoFreeConns is returned if all Client.MaxConnsPerHost connections
// to the requested host are busy.
//
// It is recommended obtaining req and resp via AcquireRequest
// and AcquireResponse in performance-critical code.
//
// Warning: ProxyTimeout does not terminate the request itself. The request will
// continue in the background and the response will be discarded.
// If requests take too long and the connection pool gets filled up please
// try setting a ReadTimeout.
func (c *Client) ProxyTimeout(addr string, req *fasthttp.Request, resp *fasthttp.Response, timeout time.Duration) error {
	// 获取fasthttp.HostClient和协议
	client, scheme, err := c.getHostClient(addr)
	if err != nil {
		return err
	}

	// 补充请求数据
	request := req
	request.URI().SetScheme(scheme)
	request.Header.ResetConnectionClose()
	request.Header.Set("Connection", "keep-alive") // 长连接

	// 检查 HTTP 响应头中是否有 "Connection: close" 标记。
	// HTTP 1.1 协议规定，如果一个 HTTP 消息的头部包含 "Connection: close" 标记，那么这个 HTTP 连接在完成当前请求后会被关闭，而不会被保持开启等待下一个请求。这个标记通常在服务器无法或不愿保持连接开启时使用。
	// connectionClose true 表示接收到的HTTP响应要求关闭连接， false则表示连接可以被保持开启。
	connectionClose := resp.ConnectionClose()

	// 转发请求
	err = client.DoTimeout(request, resp, timeout)
	if err != nil {
		return err
	}

	// 设置 HTTP 响应头的 "Connection" 字段为 "close"
	if connectionClose {
		resp.SetConnectionClose()
	}
	return nil
}

// mCleaner 清理m，如果m中的某个HostClient没有连接，那么将其从m中移除。如果m为空，那么停止清理。
func (c *Client) mCleaner(m map[string]*fasthttp.HostClient) {
	mustStop := false
	sleep := time.Second * 10
	for {
		c.mLock.Lock()
		for k, v := range m {
			// 前活动的连接数
			shouldRemove := v.ConnsCount() == 0
			if shouldRemove {
				delete(m, k)
			}
		}
		if len(m) == 0 {
			mustStop = true
		}

		c.mLock.Unlock()

		if mustStop {
			break
		}
		time.Sleep(sleep)
	}
}

// readAddress 读取 1、http协议 2、host:ip
// eg: http://demo.apinto.com:8280
func readAddress(addr string) (scheme, host string) {
	// 根据 :// 切割
	if i := strings.Index(addr, "://"); i > 0 {
		return strings.ToLower(addr[:i]), addr[i+3:]
	}
	return "http", addr
}

func addMissingPort(addr string, isTLS bool) string {
	n := strings.Index(addr, ":")
	if n >= 0 {
		return addr
	}
	port := 80
	if isTLS {
		port = 443
	}
	return net.JoinHostPort(addr, strconv.Itoa(port))
}
