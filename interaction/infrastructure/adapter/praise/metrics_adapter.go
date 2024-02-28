// Package adapter 端口-适配器模式的 adapter
package adapter

// MetricsAdapter Metrics 的 Adapter
type MetricsAdapter struct{}

// CounterIncr 监控 Counter
func (a *MetricsAdapter) CounterIncr(name string) {

}
