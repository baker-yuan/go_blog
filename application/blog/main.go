package main

import (
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func init() {
	// todo 临时设置
	trpc.ServerConfigPath = "/Users/yuanyu/code/go-study/go-blog/go_blog/blog/trpc_go.yaml"
}

func main() {
	s := trpc.NewServer()
	if err := s.Serve(); err != nil {
		log.Errorf("sever start fail: %+v", err)
	}
}
