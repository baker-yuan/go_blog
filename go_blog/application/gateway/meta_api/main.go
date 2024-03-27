package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/conf"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/server"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/storage"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/core/store"
	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
)

//func init() {
//	// 配置文件
//	rootCmd.PersistentFlags().StringVarP(&conf.ConfigFile, "config", "c", "", "config file")
//	// 工作目录
//	rootCmd.PersistentFlags().StringVarP(&conf.WorkDir, "work-dir", "p", ".", "current work directory")
//
//	//
//	rootCmd.AddCommand(
//		newVersionCommand(),
//	)
//}

func init() {
	//conf.ConfigFile
	//conf.WorkDir
}

func main() {
	// 配置初始化
	conf.InitConf()
	// 日志初始化
	log.InitLogger()
	// 创建服务
	s, err := server.NewServer(&server.Options{})
	if err != nil {
		panic(err)
	}
	// 启动服务
	// start Manager API server
	errSig := make(chan error, 5)
	s.Start(errSig)

	// start etcd connection checker
	stopEtcdConnectionChecker := etcdConnectionChecker()

	// 监听关闭信号
	// Signal received to the process externally.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-quit:
		log.Infof("The Manager API server receive %s and start shutting down", sig.String())
		stopEtcdConnectionChecker()
		s.Stop()
		log.Infof("See you next time!")
	case err := <-errSig:
		log.Errorf("The Manager API server start failed: %s", err.Error())
	}
}

func etcdConnectionChecker() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.TODO())
	unavailableTimes := 0

	go func() {
		etcdClient := storage.GenEtcdStorage().GetClient()
		for {
			select {
			case <-time.Tick(10 * time.Second):
				sCtx, sCancel := context.WithTimeout(ctx, 5*time.Second)
				err := etcdClient.Sync(sCtx)
				sCancel()
				if err != nil {
					unavailableTimes++
					log.Errorf("etcd connection loss detected, times: %d", unavailableTimes)
					continue
				}

				// After multiple failures, the connection is restored
				if unavailableTimes >= 1 {
					log.Warnf("etcd connection recovered, but after several connection losses, reinitializing stores, times: %d", unavailableTimes)
					unavailableTimes = 0

					// When this happens, force a full re-initialization of the store
					store.RangeStore(func(key store.HubKey, store *store.GenericStore) bool {
						log.Warnf("etcd store reinitializing: resource: %s", key)
						if err := store.Init(); err != nil {
							log.Errorf("etcd store reinitialize failed: resource: %s, error: %s", key, err)
						}
						return true
					})
				} else {
					log.Info("etcd connection is fine")
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// Timed re-initialization when etcd watch actively exits
	go func() {
		for {
			select {
			case <-time.Tick(2 * time.Minute):
				err := store.ReInit()
				if err != nil {
					log.Errorf("resource re-initialize failed, err: %v", err)
				}
			}
		}
	}()

	return cancel
}
