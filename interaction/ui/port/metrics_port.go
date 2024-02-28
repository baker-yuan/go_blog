// Package port 端口-适配器模式的 port
package port

// MetricsPort Metrics 的 Port
type MetricsPort interface {
	// CounterIncr 数据上报的 counter incr 接口
	CounterIncr(name string)
}
