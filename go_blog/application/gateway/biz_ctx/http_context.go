package biz_ctx

import (
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// IBodyDataReader 请求体读取
type IBodyDataReader interface {
	// ContentType "Content-Type" 字段用于表示发送给接收者的实体主体的媒体类型。
	ContentType() string
	// BodyForm 根据 "Content-Type" 头部字段解析请求体。
	// 它支持 "application/x-www-form-urlencoded" 和 "multipart/form-data" 类型的内容。
	// 与原生 request.Form 不同，此方法不包括查询参数。
	BodyForm() (url.Values, error)
	// Files 当 "Content-Type" 为 "multipart/form-data" 时有效。
	// 返回一个包含文件头信息的映射，映射的键是表单中的文件字段名称。
	Files() (map[string][]*multipart.FileHeader, error)
	// GetForm 根据给定的表单字段名称（key）返回请求体中相应的值。
	// 如果表单中不存在该字段，则返回空字符串。
	GetForm(key string) string
	// GetFile 根据给定的表单文件字段名称（key）返回请求体中相应的文件头信息。
	// 如果表单中不存在该文件字段，则返回 nil 和 false。
	GetFile(key string) (file []*multipart.FileHeader, has bool)
	// RawBody 返回请求体的原始字节数据。
	RawBody() ([]byte, error)
}

// IHeaderReader 请求头读取
type IHeaderReader interface {
	// RawHeader 返回请求头的原始字符串表示形式，包括所有的字段和它们的值。
	// 这通常用于调试或日志记录，以查看完整的请求头内容。
	RawHeader() string
	// GetHeader 根据给定的字段名称（name）返回请求头中相应的值。
	// 如果请求头中不存在该字段，则返回空字符串。
	GetHeader(name string) string
	// Headers 返回一个http.Header类型的映射，它包含了请求头中所有的字段和值。
	// map[string][]string key=请求头的字段名称 value=该字段对应的一个或多个值
	Headers() http.Header
	// Host 返回 HTTP 请求头中的 "Host" 字段的值。
	// "Host" 字段是 HTTP/1.1 请求中必须存在的字段，它指定了请求的目标主机名和（可选的）端口号。
	// 例如，"example.com" 或 "example.com:8080"。
	Host() string
	// GetCookie 根据给定的 cookie 名称（key）返回请求头中相应的 cookie 值。
	// 如果请求头中不存在该 cookie，则返回空字符串。
	GetCookie(key string) string
}

// IQueryReader url查询参数获取
type IQueryReader interface {
	// GetQuery 根据key获取查询参数
	GetQuery(key string) string
	// RawQuery 获取所有的查询参数，url ? 后面的
	RawQuery() string
}

// IURIReader 提供了一组方法，用于从 URI 中提取和读取各个组成部分的信息。
// 标准的 URL 格式遵循以下模式：
// Scheme://Host/Path?Query#Fragment
// 其中各部分的含义如下：
// - Scheme: 访问资源所使用的协议类型，例如 "http" 或 "https"。
// - Host: 资源所在服务器的地址，可能包括端口号，例如 "example.com:8080"。
// - Path: 服务器上资源的具体路径，例如 "/path/to/resource"。
// - Query: 服务器用于进一步处理请求的额外参数，例如 "query=123&name=abc"。
// - Fragment: 页面内部的锚点，不会发送到服务器，例如 "section1"。
// 例如，对于完整的 URL "http://example.com/path?query=123#fragment"：
// - RequestURI() 返回路径和查询字符串部分，即 "/path?query=123"。
// - Scheme() 返回协议方案，即 "http"。
// - RawURL() 返回不包含 fragment 的完整 URL，即 "http://example.com/path?query=123"。
// - Host() 返回主机名，即 "example.com"。
// - Path() 返回路径部分，即 "/path"。
type IURIReader interface {
	IQueryReader        // 获取url查询参数
	RequestURI() string // 返回请求的原始 URI，包括路径和查询字符串。
	Scheme() string     // 返回请求使用的协议方案。
	RawURL() string     // 返回完整的请求 URI，包括协议、主机名和路径。
	Host() string       // 返回请求的主机名，可能包括端口号。
	Path() string       // 返回请求的路径部分。
}

// IRequestReader 请求数据读取接口
type IRequestReader interface {
	URI() IURIReader       // url读取
	Header() IHeaderReader // 请求头读取
	Body() IBodyDataReader // 请求体读取

	// RemoteAddr 获取客户端地址
	// 直接与服务器建立连接的客户端的网络地址。如果请求没有经过任何代理，那么这个地址就是发起请求的客户端的真实 IP 地址和端口号。
	// 然而，如果请求是通过一个或多个代理（如 HTTP 代理、负载均衡器、CDN 等）转发的，返回最后一个代理的地址，而不是原始客户端的地址。这是因为在TCP层面上，服务器看到的是与其直接建立连接的实体的地址。
	// 在处理经过代理的请求时，如果您想获取原始客户端的IP地址，应该查看X-Forwarded-For或X-Real-IP这样的 HTTP 头部。这些头部由代理添加，用于传递原始客户端的IP地址信息。但是请注意，这些头部可能被恶意用户篡改，因此在信任这些头部之前，应该确保它们来自可信的代理。
	RemoteAddr() string // 客户端ip地址
	RemotePort() string // 客户端端口

	// RealIP 获取真实IP
	RealIP() string    // 客户端ip
	ForwardIP() string // X-Forwarded-For HTTP请求头用于记录整个请求链路中所有经过的代理服务器的IP地址。当客户端通过一个或多个代理服务器发送请求时，每个代理服务器都会在X-Forwarded-For头部中追加自己的IP地址。

	Method() string      // HTTP请求的方法，如 GET、POST、PUT、DELETE 等
	ContentLength() int  // HTTP请求中Content-Length头部的值，它表示HTTP消息正文的长度，单位是字节
	ContentType() string // HTTP请求Content-Type头部，它描述了HTTP消息正文的媒体类型（也称为 MIME 类型）
	String() string      // 整个HTTP请求的详细信息，包括请求行（如方法、URI和HTTP版本），请求头，以及请求体（如果有的话）。这个字符串表示形式主要用于调试目的，因为它可以让你看到完整的请求内容。
}

// IHttpContext 扩展GatewayContext接口，定义http协议特有的
type IHttpContext interface {
	// IBizContext 组合IBizContext
	IBizContext
	// Request 请求数据读取接口
	Request() IRequestReader

	// SendTo 如果下游是http服务，通过这个方法转发到下游
	SendTo(scheme string, node IInstance, timeout time.Duration) error
	// FastFinish 结束请求，释放资源
	FastFinish()
}
