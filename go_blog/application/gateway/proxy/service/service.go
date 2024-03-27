package service

import (
	"fmt"

	"github.com/baker-yuan/go-blog/application/blog/gateway/proxy/biz_ctx"
	"github.com/polarismesh/polaris-go"
)

var consumer polaris.ConsumerAPI

func Init() error {
	c, err := polaris.NewConsumerAPI()
	if err != nil {
		return err
	}
	consumer = c
	return nil
}

func GetOneInstance(namespace string, service string) (string, error) {
	getOneRequest := &polaris.GetOneInstanceRequest{}
	getOneRequest.Namespace = namespace
	getOneRequest.Service = service
	oneInstResp, err := consumer.GetOneInstance(getOneRequest)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", oneInstResp.GetInstance().GetHost(), oneInstResp.GetInstance().GetPort()), nil
}

// 强制resourceRepo实现biz.ResourceRepo
var _ biz_ctx.IInstance = &InstanceImpl{}

type InstanceImpl struct {
}

func (i InstanceImpl) GetAttrs() biz_ctx.Attrs {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) GetAttrByName(name string) (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) ID() string {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) IP() string {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) Port() int {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) Addr() string {
	return "127.0.0.1:9003"
}

func (i InstanceImpl) Status() biz_ctx.NodeStatus {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) Up() {
	//TODO implement me
	panic("implement me")
}

func (i InstanceImpl) Down() {
	//TODO implement me
}
