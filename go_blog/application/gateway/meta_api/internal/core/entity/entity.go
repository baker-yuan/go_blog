package entity

import (
	"reflect"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils"
)

// BaseInfo 公共属性
type BaseInfo struct {
	ID         interface{} `json:"id"`                    // 分布式唯一ID
	CreateTime int64       `json:"create_time,omitempty"` // 创建时间
	UpdateTime int64       `json:"update_time,omitempty"` // 更新时间
}

func (info *BaseInfo) GetBaseInfo() *BaseInfo {
	return info
}

func (info *BaseInfo) Creating() {
	if info.ID == nil {
		info.ID = utils.GetFlakeUidStr()
	} else {
		// convert to string if it's not
		if reflect.TypeOf(info.ID).String() != "string" {
			info.ID = utils.InterfaceToString(info.ID)
		}
	}
	info.CreateTime = time.Now().Unix()
	info.UpdateTime = time.Now().Unix()
}

func (info *BaseInfo) Updating(storedInfo *BaseInfo) {
	info.ID = storedInfo.ID
	info.CreateTime = storedInfo.CreateTime
	info.UpdateTime = time.Now().Unix()
}

func (info *BaseInfo) KeyCompat(key string) {
	if info.ID == nil && key != "" {
		info.ID = key
	}
}

// Status 发布状态
type Status uint8

// swagger:model 路由
type Route struct {
	BaseInfo
	// 基本信息
	Name            string            `json:"name"`                       // 路由名称
	Labels          map[string]string `json:"labels,omitempty"`           // 标签(为路由增加自定义标签，可用于路由分组。)
	Desc            string            `json:"desc,omitempty"`             // 路由描述信息
	ServiceID       interface{}       `json:"service_id,omitempty"`       // 需要绑定的 Service id
	EnableWebsocket bool              `json:"enable_websocket,omitempty"` // 允许WebSocket
	Status          Status            `json:"status"`                     // 发布状态

	// 匹配条件
	Host        string   `json:"host,omitempty"`         // 域名(路由匹配的域名列表。支持泛域名，如：*.test.com)
	Hosts       []string `json:"hosts,omitempty"`        //
	URI         string   `json:"uri,omitempty"`          // 路径(HTTP 请求路径，如 /foo/index.html，支持请求路径前缀 /foo/*。/* 代表所有路径)
	Uris        []string `json:"uris,omitempty"`         //
	RemoteAddr  string   `json:"remote_addr,omitempty"`  // 客户端地址(客户端与服务器握手时 IP，即客户端 IP，例如：192.168.1.101，192.168.1.0/24，::1，fe80::1，fe80::1/64)
	RemoteAddrs []string `json:"remote_addrs,omitempty"` //
	Methods     []string `json:"methods,omitempty"`      // HTTP 方法
	Priority    int      `json:"priority,omitempty"`     // 优先级

	// 高级匹配条件
	Vars []interface{} `json:"vars,omitempty"` //

	// 插件
	Plugins        map[string]interface{} `json:"plugins,omitempty"`          //
	PluginConfigID interface{}            `json:"plugin_config_id,omitempty"` //

	// 上游
	Upstream   *UpstreamDef `json:"upstream,omitempty"`    // 手动填写
	UpstreamID interface{}  `json:"upstream_id,omitempty"` // 选择的上游ID

	// 其他
	FilterFunc      string      `json:"filter_func,omitempty"`      //
	Script          interface{} `json:"script,omitempty"`           //
	ScriptID        interface{} `json:"script_id,omitempty"`        // For debug and optimization(cache), currently same as Route's ID
	ServiceProtocol string      `json:"service_protocol,omitempty"` //
}

// --- structures for upstream start  ---

// TimeoutValue 超时
type TimeoutValue float32

// Timeout 超时
type Timeout struct {
	Connect TimeoutValue `json:"connect,omitempty"` // 连接超时(建立从请求到上游服务器的连接的超时时间)
	Send    TimeoutValue `json:"send,omitempty"`    // 发送超时(发送数据到上游服务器的超时时间)
	Read    TimeoutValue `json:"read,omitempty"`    // 接收超时(从上游服务器接收数据的超时时间)
}

// Node 目标节点
type Node struct {
	Host     string      `json:"host,omitempty"`     // 主机名
	Port     int         `json:"port,omitempty"`     // 端口
	Weight   int         `json:"weight"`             // 权重
	Metadata interface{} `json:"metadata,omitempty"` // 元数据
	Priority int         `json:"priority,omitempty"` // 优先级
}

type K8sInfo struct {
	Namespace   string `json:"namespace,omitempty"`
	DeployName  string `json:"deploy_name,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	Port        int    `json:"port,omitempty"`
	BackendType string `json:"backend_type,omitempty"`
}

// Healthy 健康状态
type Healthy struct {
	Interval     int   `json:"interval,omitempty"`      // 间隔时间(对健康的上游服务目标节点进行主动健康检查的间隔时间（以秒为单位）。数值为0表示对健康节点不进行主动健康检查。)
	Successes    int   `json:"successes,omitempty"`     // 成功次数(主动健康检查的 HTTP 成功次数，若达到此值，表示上游服务目标节点是健康的。)
	HttpStatuses []int `json:"http_statuses,omitempty"` // 状态码(HTTP 状态码列表，当探针在主动健康检查中返回时，视为健康。)
}

// UnHealthy 不健康状态
type UnHealthy struct {
	Timeouts     int   `json:"timeouts,omitempty"`      // 超时时间(活动探针中认为目标不健康的超时次数。)
	Interval     int   `json:"interval,omitempty"`      // 间隔时间(对不健康目标的主动健康检查之间的间隔（以秒为单位）。数值为0表示不应该对健康目标进行主动探查。)
	HTTPStatuses []int `json:"http_statuses,omitempty"` // 状态码
	TCPFailures  int   `json:"tcp_failures,omitempty"`  // HTTP失败次数(主动健康检查的 HTTP 失败次数，默认值为0。若达到此值，表示上游服务目标节点是不健康的。)
	HTTPFailures int   `json:"http_failures,omitempty"` // TCP失败次数(主动探测中 TCP 失败次数超过该值时，认为目标不健康。)
}

// Active 主动检查
type Active struct {
	Type                   string       `json:"type,omitempty"`                     // 类型(是使用 HTTP 或 HTTPS 进行主动健康检查，还是只尝试 TCP 连接。)
	Timeout                TimeoutValue `json:"timeout,omitempty"`                  // 超时时间(主动健康检查的套接字的超时时间)
	Concurrency            int          `json:"concurrency,omitempty"`              // 并行数量(在主动健康检查中同时检查的目标数量。)
	Host                   string       `json:"host,omitempty"`                     // 主机名(进行主动健康检查时使用的 HTTP 请求主机名)
	Port                   int          `json:"port,omitempty"`                     // 端口
	HTTPPath               string       `json:"http_path,omitempty"`                // 请求路径(向目标节点发出 HTTP GET 请求时应使用的路径。)
	HTTPSVerifyCertificate bool         `json:"https_verify_certificate,omitempty"` //
	ReqHeaders             []string     `json:"req_headers,omitempty"`              // 请求头
	Healthy                Healthy      `json:"healthy,omitempty"`                  // 健康状态
	UnHealthy              UnHealthy    `json:"unhealthy,omitempty"`                // 不健康状态
}

// Passive 被动检查(启用被动健康检查时，需要同时启用主动健康检查。)
type Passive struct {
	Type      string    `json:"type,omitempty"`      // 类型(是使用 HTTP 或 HTTPS 进行主动健康检查，还是只尝试 TCP 连接。)
	Healthy   Healthy   `json:"healthy,omitempty"`   // 健康状态
	UnHealthy UnHealthy `json:"unhealthy,omitempty"` // 不健康状态
}

// HealthChecker 健康检查
type HealthChecker struct {
	Active  Active  `json:"active,omitempty"`  // 主动检查
	Passive Passive `json:"passive,omitempty"` // 被动检查
}

// UpstreamTLS 证书
type UpstreamTLS struct {
	ClientCert string `json:"client_cert,omitempty"`
	ClientKey  string `json:"client_key,omitempty"`
}

// UpstreamKeepalivePool 连接池(为 upstream 对象设置独立的连接池)
type UpstreamKeepalivePool struct {
	IdleTimeout *TimeoutValue `json:"idle_timeout,omitempty"` // 容量
	Requests    int           `json:"requests,omitempty"`     // 空闲超时时间
	Size        int           `json:"size"`                   // 请求数量
}

// UpstreamDef 上游服务
type UpstreamDef struct {
	Name          string                 `json:"name,omitempty"`           // 名称
	Desc          string                 `json:"desc,omitempty"`           // 描述
	Type          string                 `json:"type,omitempty"`           // 负载均衡算法
	DiscoveryType string                 `json:"discovery_type,omitempty"` // 上游类型 节点、服务发现
	ServiceName   string                 `json:"service_name,omitempty"`   //
	DiscoveryArgs map[string]string      `json:"discovery_args,omitempty"` //
	Nodes         interface{}            `json:"nodes,omitempty"`          //
	Retries       *int                   `json:"retries,omitempty"`        // 重试次数(重试机制将请求发到下一个上游节点。值为 0 表示禁用重试机制，留空表示使用可用后端节点的数量。)
	RetryTimeout  TimeoutValue           `json:"retry_timeout,omitempty"`  // 重试超时时间(限制是否继续重试的时间，若之前的请求和重试请求花费太多时间就不再继续重试。0 代表不启用重试超时机制。)
	Scheme        string                 `json:"scheme,omitempty"`         // 协议
	Timeout       *Timeout               `json:"timeout,omitempty"`        // 超时
	KeepalivePool *UpstreamKeepalivePool `json:"keepalive_pool,omitempty"` // 连接池
	Checks        interface{}            `json:"checks,omitempty"`         // 健康检查
	HashOn        string                 `json:"hash_on,omitempty"`        //
	Key           string                 `json:"key,omitempty"`            //
	PassHost      string                 `json:"pass_host,omitempty"`      //
	UpstreamHost  string                 `json:"upstream_host,omitempty"`  //
	Labels        map[string]string      `json:"labels,omitempty"`         //
	TLS           *UpstreamTLS           `json:"tls,omitempty"`            //
}

// swagger:model 上游
type Upstream struct {
	BaseInfo
	UpstreamDef
}

type UpstreamNameResponse struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}

func (upstream *Upstream) Parse2NameResponse() (*UpstreamNameResponse, error) {
	nameResp := &UpstreamNameResponse{
		ID:   upstream.ID,
		Name: upstream.Name,
	}
	return nameResp, nil
}

// --- structures for upstream end  ---

// swagger:model Consumer
type Consumer struct {
	Username   string                 `json:"username"`              // 名称
	Desc       string                 `json:"desc,omitempty"`        // 描述
	Plugins    map[string]interface{} `json:"plugins,omitempty"`     // 插件
	Labels     map[string]string      `json:"labels,omitempty"`      //
	CreateTime int64                  `json:"create_time,omitempty"` // 创建时间
	UpdateTime int64                  `json:"update_time,omitempty"` // 更新时间
}

type SSLClient struct {
	CA    string `json:"ca,omitempty"`
	Depth int    `json:"depth,omitempty"`
}

// swagger:model SSL
type SSL struct {
	BaseInfo
	Cert          string            `json:"cert,omitempty"`
	Key           string            `json:"key,omitempty"`
	Sni           string            `json:"sni,omitempty"`
	Snis          []string          `json:"snis,omitempty"`
	Certs         []string          `json:"certs,omitempty"`
	Keys          []string          `json:"keys,omitempty"`
	ExpTime       int64             `json:"exptime,omitempty"`
	Status        int               `json:"status"`
	ValidityStart int64             `json:"validity_start,omitempty"`
	ValidityEnd   int64             `json:"validity_end,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Client        *SSLClient        `json:"client,omitempty"`
}

// swagger:model 服务
type Service struct {
	BaseInfo
	Name            string                 `json:"name,omitempty"`             // 名称
	Desc            string                 `json:"desc,omitempty"`             // 描述
	Hosts           []string               `json:"hosts,omitempty"`            // 域名(路由匹配的域名列表。支持泛域名，如：*.test.com)
	Upstream        *UpstreamDef           `json:"upstream,omitempty"`         // 直接录入上游
	UpstreamID      interface{}            `json:"upstream_id,omitempty"`      // 选择的上游ID
	Plugins         map[string]interface{} `json:"plugins,omitempty"`          // 插件
	Script          string                 `json:"script,omitempty"`           //
	Labels          map[string]string      `json:"labels,omitempty"`           //
	EnableWebsocket bool                   `json:"enable_websocket,omitempty"` //
}

type Script struct {
	ID     string      `json:"id"`
	Script interface{} `json:"script,omitempty"`
}

type RequestValidation struct {
	Type       string      `json:"type,omitempty"`
	Required   []string    `json:"required,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
}

// swagger:model 全局插件
type GlobalPlugins struct {
	BaseInfo
	Plugins map[string]interface{} `json:"plugins"`
}

type ServerInfo struct {
	BaseInfo
	LastReportTime int64  `json:"last_report_time,omitempty"`
	UpTime         int64  `json:"up_time,omitempty"`
	BootTime       int64  `json:"boot_time,omitempty"`
	EtcdVersion    string `json:"etcd_version,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	Version        string `json:"version,omitempty"`
}

// swagger:model GlobalPlugins
type PluginConfig struct {
	BaseInfo
	Desc    string                 `json:"desc,omitempty"`
	Plugins map[string]interface{} `json:"plugins"`
	Labels  map[string]string      `json:"labels,omitempty"`
}

// swagger:model Proto文件
type Proto struct {
	BaseInfo        //
	Desc     string `json:"desc,omitempty"` // 描述
	Content  string `json:"content"`        // 内容
}

// swagger:model StreamRoute
// https://apisix.apache.org/blog/2022/07/29/release-apache-apisix-2.15/#allow-collection-of-metrics-on-stream-route
// https://blog.csdn.net/weixin_44917045/article/details/131835341
type StreamRoute struct {
	BaseInfo
	Desc       string                 `json:"desc,omitempty"`        // 描述信息，用于描述这个StreamRoute的用途或其他信息。
	RemoteAddr string                 `json:"remote_addr,omitempty"` // 远程地址，用于匹配客户端的IP地址或CIDR范围。
	ServerAddr string                 `json:"server_addr,omitempty"` // 服务器地址，用于指定APISIX监听的地址。
	ServerPort int                    `json:"server_port,omitempty"` // 服务器端口，用于指定APISIX监听的端口。
	SNI        string                 `json:"sni,omitempty"`         // 用于TLS连接的Server Name Indication，用于匹配客户端的SNI。
	Upstream   *UpstreamDef           `json:"upstream,omitempty"`    // 上游服务的定义，包括上游节点的地址和端口，以及负载均衡策略等。
	UpstreamID interface{}            `json:"upstream_id,omitempty"` // 上游服务的ID，可以用来引用一个已经定义的上游服务。
	Plugins    map[string]interface{} `json:"plugins,omitempty"`     // 插件的定义，用于在这个StreamRoute上启用一些插件，如限流插件、认证插件等。
}

// swagger:model SystemConfig
type SystemConfig struct {
	ConfigName string                 `json:"config_name"`
	Desc       string                 `json:"desc,omitempty"`
	Payload    map[string]interface{} `json:"payload,omitempty"`
	CreateTime int64                  `json:"create_time,omitempty"`
	UpdateTime int64                  `json:"update_time,omitempty"`
}
