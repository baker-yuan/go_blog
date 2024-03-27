package entity

type GetBaseInfo interface {
	GetBaseInfo() *BaseInfo
}

type GetPlugins interface {
	GetPlugins() map[string]interface{}
}

func (r *Route) GetPlugins() map[string]interface{} {
	return r.Plugins
}

func (s *Service) GetPlugins() map[string]interface{} {
	return s.Plugins
}

func (c *Consumer) GetPlugins() map[string]interface{} {
	return c.Plugins
}

func (g *GlobalPlugins) GetPlugins() map[string]interface{} {
	return g.Plugins
}

func (p *PluginConfig) GetPlugins() map[string]interface{} {
	return p.Plugins
}
