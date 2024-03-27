package store

import (
	"fmt"
	"reflect"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/entity"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/utils"
)

type HubKey string

// https://apisix.apache.org/zh/docs/apisix/terminology/api-gateway/
const (
	HubKeyConsumer     HubKey = "consumer"
	HubKeyRoute        HubKey = "route"
	HubKeyService      HubKey = "service"
	HubKeySsl          HubKey = "ssl"
	HubKeyUpstream     HubKey = "upstream"
	HubKeyScript       HubKey = "script"
	HubKeyGlobalRule   HubKey = "global_rule"
	HubKeyServerInfo   HubKey = "server_info"
	HubKeyPluginConfig HubKey = "plugin_config"
	HubKeyProto        HubKey = "proto"
	HubKeyStreamRoute  HubKey = "stream_route"
	HubKeySystemConfig HubKey = "system_config"
)

var (
	storeHub = map[HubKey]*GenericStore{}
)

func InitStore(key HubKey, opt GenericStoreOption) error {
	// 那些资源需要检查
	hubsNeedCheck := map[HubKey]bool{
		HubKeyConsumer:     true,
		HubKeyRoute:        true,
		HubKeySsl:          true,
		HubKeyService:      true,
		HubKeyUpstream:     true,
		HubKeyGlobalRule:   true,
		HubKeyStreamRoute:  true,
		HubKeySystemConfig: true,
	}

	if _, ok := hubsNeedCheck[key]; ok {
		validator, err := NewAPISIXJsonSchemaValidator("main." + string(key))
		if err != nil {
			return err
		}
		opt.Validator = validator
	}
	opt.HubKey = key
	s, err := NewGenericStore(opt)
	if err != nil {
		log.Errorf("NewGenericStore error: %s", err)
		return err
	}
	if err := s.Init(); err != nil {
		log.Errorf("GenericStore init error: %s", err)
		return err
	}

	utils.AppendToClosers(s.Close)
	storeHub[key] = s
	return nil
}

func GetStore(key HubKey) *GenericStore {
	if s, ok := storeHub[key]; ok {
		return s
	}
	panic(fmt.Sprintf("no store with key: %s", key))
}

func RangeStore(f func(key HubKey, store *GenericStore) bool) {
	for k, s := range storeHub {
		if k != "" && s != nil {
			if !f(k, s) {
				break
			}
		}
	}
}

func InitStores() error {
	// 消费者
	err := InitStore(HubKeyConsumer, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/consumers",
		ObjType:  reflect.TypeOf(entity.Consumer{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.Consumer)
			return r.Username
		},
	})
	if err != nil {
		return err
	}

	// 路由
	err = InitStore(HubKeyRoute, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/routes",
		ObjType:  reflect.TypeOf(entity.Route{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.Route)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	// 服务
	err = InitStore(HubKeyService, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/services",
		ObjType:  reflect.TypeOf(entity.Service{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.Service)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	// 证书
	err = InitStore(HubKeySsl, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/ssls",
		ObjType:  reflect.TypeOf(entity.SSL{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.SSL)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	// 上游
	err = InitStore(HubKeyUpstream, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/upstreams",
		ObjType:  reflect.TypeOf(entity.Upstream{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.Upstream)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	// 脚本
	err = InitStore(HubKeyScript, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/scripts",
		ObjType:  reflect.TypeOf(entity.Script{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.Script)
			return r.ID
		},
	})
	if err != nil {
		return err
	}

	// 全局插件配置
	err = InitStore(HubKeyGlobalRule, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/global_rules",
		ObjType:  reflect.TypeOf(entity.GlobalPlugins{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.GlobalPlugins)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	err = InitStore(HubKeyServerInfo, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/data_plane/server_info",
		ObjType:  reflect.TypeOf(entity.ServerInfo{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.ServerInfo)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	err = InitStore(HubKeyPluginConfig, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/plugin_configs",
		ObjType:  reflect.TypeOf(entity.PluginConfig{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.PluginConfig)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	err = InitStore(HubKeyProto, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/protos",
		ObjType:  reflect.TypeOf(entity.Proto{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.Proto)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	err = InitStore(HubKeyStreamRoute, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/stream_routes",
		ObjType:  reflect.TypeOf(entity.StreamRoute{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.StreamRoute)
			return utils.InterfaceToString(r.ID)
		},
	})
	if err != nil {
		return err
	}

	err = InitStore(HubKeySystemConfig, GenericStoreOption{
		BasePath: conf.ETCDConfig.Prefix + "/system_config",
		ObjType:  reflect.TypeOf(entity.SystemConfig{}),
		KeyFunc: func(obj interface{}) string {
			r := obj.(*entity.SystemConfig)
			return r.ConfigName
		},
	})
	if err != nil {
		return err
	}

	return nil
}
