// Package server http和grpc实例的创建和配置
package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewTRPCServer)
