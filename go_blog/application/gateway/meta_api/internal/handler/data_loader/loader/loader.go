package loader

import "github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/entity"

// DataSets are intermediate structures used to handle
// import and export data with APISIX entities.
// On import, raw data will be parsed as DataSets
// On export, DataSets will be encoded to raw data
type DataSets struct {
	Routes        []entity.Route
	Upstreams     []entity.Upstream
	Services      []entity.Service
	Consumers     []entity.Consumer
	SSLs          []entity.SSL
	StreamRoutes  []entity.StreamRoute
	GlobalPlugins []entity.GlobalPlugins
	PluginConfigs []entity.PluginConfig
	Protos        []entity.Proto
}

// Loader provide data loader abstraction
type Loader interface {
	// Import accepts data and converts it into entity data sets
	Import(input interface{}) (*DataSets, error)

	// Export accepts entity data sets and converts it into a specific format
	Export(data DataSets) (interface{}, error)
}
