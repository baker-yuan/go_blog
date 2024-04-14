package server

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/meta_api/internal/core/storage"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/meta_api/internal/core/store"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/gateway/meta_api/internal/log"
)

func (s *server) setupStore() error {
	// 初始化etcd客户端
	if err := storage.InitETCDClient(conf.ETCDConfig); err != nil {
		log.Errorf("init etcd client fail: %v", err)
		return err
	}
	// 初始化每个具体的资源类型操作etcd
	if err := store.InitStores(); err != nil {
		log.Errorf("init stores fail: %v", err)
		return err
	}
	return nil
}
