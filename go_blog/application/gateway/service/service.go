package service

import (
	"fmt"

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
