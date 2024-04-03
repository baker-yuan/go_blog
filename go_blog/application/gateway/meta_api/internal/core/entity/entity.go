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

// GetBaseInfo 获取公共属性
func (info *BaseInfo) GetBaseInfo() *BaseInfo {
	return info
}

// Creating 初始化公共属性
func (info *BaseInfo) Creating() {
	if info.ID == nil {
		// ID为nil设置一个分布式ID
		info.ID = utils.GetFlakeUidStr()
	} else {
		// 确保ID是string类型的
		// convert to string if it's not
		if reflect.TypeOf(info.ID).String() != "string" {
			info.ID = utils.InterfaceToString(info.ID)
		}
	}
	info.CreateTime = time.Now().Unix()
	info.UpdateTime = time.Now().Unix()
}

// Updating 修改公共属性
func (info *BaseInfo) Updating(storedInfo *BaseInfo) {
	info.ID = storedInfo.ID
	info.CreateTime = storedInfo.CreateTime
	info.UpdateTime = time.Now().Unix()
}

// KeyCompat 比较公共属性里面的ID是否一样
func (info *BaseInfo) KeyCompat(key string) {
	if info.ID == nil && key != "" {
		info.ID = key
	}
}

// swagger:model PluginConfig
//
// 在很多情况下，我们在不同的路由中会使用相同的插件规则，此时就可以通过 Plugin Config 来设置这些规则。Plugin Config 属于一组通用插件配置的抽象。
// plugins 的配置可以通过 Admin API /apisix/admin/plugin_configs 进行单独配置，在路由中使用 plugin_config_id 与之进行关联。
// 对于同一个插件的配置，只能有一个是有效的，优先级为 Consumer > Route > Plugin Config > Service。
//
// https://apisix.apache.org/zh/docs/apisix/terminology/plugin-config/
type PluginConfig struct {
	BaseInfo                        // 公共属性
	Desc     string                 `json:"desc,omitempty"`
	Plugins  map[string]interface{} `json:"plugins"`
	Labels   map[string]string      `json:"labels,omitempty"`
}

// swagger:model 全局插件
//
// Plugin 配置可直接绑定在 Route 上，也可以被绑定在 Service 或 Consumer 上。
// 如果你需要一个能作用于所有请求的 Plugin，可以通过 Global Rules 启用一个全局的插件配置。
// 全局规则相对于 Route、Service、Plugin Config、Consumer 中的插件配置，Global Rules 中的插件总是优先执行。
//
// https://apisix.apache.org/zh/docs/apisix/terminology/global-rule/
type GlobalPlugins struct {
	BaseInfo
	Plugins map[string]interface{} `json:"plugins"`
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

// swagger:model 上游
// Upstream（也称之为上游）是对虚拟主机抽象，即应用层服务或节点的抽象。你可以通过 Upstream 对象对多个服务节点按照配置规则进行负载均衡。
// 上游的地址信息可以直接配置到路由（或服务）中。
//
// https://apisix.apache.org/zh/docs/apisix/terminology/upstream/
type Upstream struct {
	BaseInfo
	UpstreamDef
}

// UpstreamDef 上游服务
type UpstreamDef struct {
	// 基本信息
	Name   string            `json:"name,omitempty"`   // 名称
	Desc   string            `json:"desc,omitempty"`   // 描述
	Labels map[string]string `json:"labels,omitempty"` // 标签
	// 负载均衡
	// 带权轮询（Round Robin）
	// 一致性哈希（CHash）
	// 指数加权移动平均法（EWMA）
	// 最小连接数（least_conn）
	Type string `json:"type,omitempty"` // 负载均衡算法
	// 服务发现
	DiscoveryType string            `json:"discovery_type,omitempty"` // 上游类型 节点、服务发现
	DiscoveryArgs map[string]string `json:"discovery_args,omitempty"` // 服务发现-服务发现类型
	ServiceName   string            `json:"service_name,omitempty"`   // 服务发现-服务名称
	Nodes         interface{}       `json:"nodes,omitempty"`          // 节点-目标节点Node
	// Host请求头
	// pass 保持与客户端请求一致的主机名
	// node 使用目标节点列表中的主机名或IP
	PassHost string `json:"pass_host,omitempty"` // Host请求头
	//
	Retries      *int         `json:"retries,omitempty"`       // 重试次数(重试机制将请求发到下一个上游节点。值为 0 表示禁用重试机制，留空表示使用可用后端节点的数量。)
	RetryTimeout TimeoutValue `json:"retry_timeout,omitempty"` // 重试超时时间(限制是否继续重试的时间，若之前的请求和重试请求花费太多时间就不再继续重试。0 代表不启用重试超时机制。)
	Scheme       string       `json:"scheme,omitempty"`        // 协议
	Timeout      *Timeout     `json:"timeout,omitempty"`       // 超时
	// 连接池
	KeepalivePool *UpstreamKeepalivePool `json:"keepalive_pool,omitempty"` // 连接池
	// 健康检查
	Checks interface{} `json:"checks,omitempty"` // 健康检查
	// 其他
	HashOn       string       `json:"hash_on,omitempty"`       //
	Key          string       `json:"key,omitempty"`           //
	UpstreamHost string       `json:"upstream_host,omitempty"` //
	TLS          *UpstreamTLS `json:"tls,omitempty"`           //
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

// swagger:model 服务
type Service struct {
	// 基本信息
	BaseInfo
	Name   string            `json:"name,omitempty"`   // 名称
	Desc   string            `json:"desc,omitempty"`   // 描述
	Hosts  []string          `json:"hosts,omitempty"`  // 域名(路由匹配的域名列表。支持泛域名，如：*.test.com)
	Labels map[string]string `json:"labels,omitempty"` // 标签
	// 上游信息
	Upstream   *UpstreamDef `json:"upstream,omitempty"`    // 直接录入上游
	UpstreamID interface{}  `json:"upstream_id,omitempty"` // 选择的上游ID
	// 插件
	Plugins map[string]interface{} `json:"plugins,omitempty"` // 插件
	//
	EnableWebsocket bool `json:"enable_websocket,omitempty"` //
	// 其他
	Script string `json:"script,omitempty"` // 脚本
}

// Status 发布状态
type Status uint8

// swagger:model 路由
type Route struct {
	BaseInfo // 公共属性
	// 基本信息
	Name            string            `json:"name"`                       // 路由名称
	Labels          map[string]string `json:"labels,omitempty"`           // 标签(为路由增加自定义标签，可用于路由分组。)
	Desc            string            `json:"desc,omitempty"`             // 路由描述信息
	ServiceID       interface{}       `json:"service_id,omitempty"`       // 需要绑定的服务ID(Service#ID)
	EnableWebsocket bool              `json:"enable_websocket,omitempty"` // 允许WebSocket
	Status          Status            `json:"status"`                     // 发布状态

	// 匹配条件
	Host        string        `json:"host,omitempty"`         // 域名(路由匹配的域名列表。支持泛域名，如：*.test.com)
	Hosts       []string      `json:"hosts,omitempty"`        //
	URI         string        `json:"uri,omitempty"`          // 路径(HTTP 请求路径，如 /foo/index.html，支持请求路径前缀 /foo/*。/* 代表所有路径)
	Uris        []string      `json:"uris,omitempty"`         //
	RemoteAddr  string        `json:"remote_addr,omitempty"`  // 客户端地址(客户端与服务器握手时 IP，即客户端 IP，例如：192.168.1.101，192.168.1.0/24，::1，fe80::1，fe80::1/64)
	RemoteAddrs []string      `json:"remote_addrs,omitempty"` //
	Methods     []string      `json:"methods,omitempty"`      // HTTP方法
	Priority    int           `json:"priority,omitempty"`     // 优先级
	Vars        []interface{} `json:"vars,omitempty"`         // 高级匹配条件

	// 插件
	Plugins        map[string]interface{} `json:"plugins,omitempty"`          // 路由直接绑定插件
	PluginConfigID interface{}            `json:"plugin_config_id,omitempty"` // 一组通用插件配置，插件配置共用 PluginConfig

	// 上游
	Upstream   *UpstreamDef `json:"upstream,omitempty"`    // 路由直接绑定上游
	UpstreamID interface{}  `json:"upstream_id,omitempty"` // 选择的上游ID

	// 其他
	ServiceProtocol string      `json:"service_protocol,omitempty"` //
	Script          interface{} `json:"script,omitempty"`           //
	ScriptID        interface{} `json:"script_id,omitempty"`        // For debug and optimization(cache), currently same as Route's ID
	FilterFunc      string      `json:"filter_func,omitempty"`      //
}

// swagger:model Consumer
// Consumer 是某类服务的消费者，需要与用户认证配合才可以使用。当不同的消费者请求同一个 API 时，APISIX 会根据当前请求的用户信息，对应不同的 Plugin 或 Upstream 配置。如果 Route、Service、Consumer 和 Plugin Config 都绑定了相同的插件，只有消费者的插件配置会生效。插件配置的优先级由高到低的顺序是：Consumer > Route > Plugin Config > Service。
// 对于 API 网关而言，一般情况可以通过请求域名、客户端 IP 地址等字段识别到某类请求方，然后进行插件过滤并转发请求到指定上游。但有时候该方式达不到用户需求，因此 APISIX 支持了 Consumer 对象。
//
// https://apisix.apache.org/zh/docs/apisix/terminology/consumer/
type Consumer struct {
	Username   string                 `json:"username"`              // 名称
	Desc       string                 `json:"desc,omitempty"`        // 描述
	Plugins    map[string]interface{} `json:"plugins,omitempty"`     // 插件
	Labels     map[string]string      `json:"labels,omitempty"`      // 标签
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

// Script 表示将在 HTTP 请求/响应生命周期期间执行的脚本。
// Script 配置需要绑定在路由上。
// Script 与 Plugin 不兼容，并且 Script 优先执行 Script，这意味着配置 Script 后，Route 上配置的 Plugin 将不被执行。
// 理论上，在 Script 中可以编写任意 Lua 代码，你也可以直接调用已有的插件以复用已有的代码。
// Script 也有执行阶段概念，支持 access、header_filter、body_filter 和 log 阶段。系统会在相应阶段中自动执行 Script 脚本中对应阶段的代码。
//
//	{
//	   ...
//	   "script": "local _M = {} \n function _M.access(api_ctx) \n ngx.log(ngx.INFO,\"hit access phase\") \n end \nreturn _M"
//	}
//
// https://apisix.apache.org/zh/docs/apisix/terminology/script/
type Script struct {
	ID     string      `json:"id"`
	Script interface{} `json:"script,omitempty"`
}

type RequestValidation struct {
	Type       string      `json:"type,omitempty"`
	Required   []string    `json:"required,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
}

// ServerInfo server-info插件可以定期将服务基本信息上报至 etcd。
//
// https://apisix.apache.org/zh/docs/apisix/plugins/server-info/
type ServerInfo struct {
	BaseInfo              // BaseInfo#ID APISIX服务实例id。
	LastReportTime int64  `json:"last_report_time,omitempty"` //
	UpTime         int64  `json:"up_time,omitempty"`          //
	BootTime       int64  `json:"boot_time,omitempty"`        // APISIX 服务实例的启动时间（UNIX 时间戳），如果对 APISIX 进行热更新操作，该值将被重置。普通的 reload 操作不会影响该值。
	EtcdVersion    string `json:"etcd_version,omitempty"`     // etcd 集群的版本信息，如果 APISIX 和 etcd 集群之间存在网络分区，该值将设置为 "unknown"。
	Hostname       string `json:"hostname,omitempty"`         // 部署 APISIX 的主机或 Pod 的主机名信息。
	Version        string `json:"version,omitempty"`          // APISIX 版本信息。
}

// swagger:model Proto文件
type Proto struct {
	BaseInfo        // 公共属性
	Desc     string `json:"desc,omitempty"` // 描述
	Content  string `json:"content"`        // 内容
}

// swagger:model StreamRoute
// https://apisix.apache.org/blog/2022/07/29/release-apache-apisix-2.15/#allow-collection-of-metrics-on-stream-route
// https://blog.csdn.net/weixin_44917045/article/details/131835341
type StreamRoute struct {
	BaseInfo                          // 公共属性
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
