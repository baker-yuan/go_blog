package entity

import (
	"errors"
	"net"
	"strconv"
	"strings"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/internal/log"
)

func mapKV2Node(key string, val float64) (*Node, error) {
	host, port, err := net.SplitHostPort(key)

	// ipv6 address
	if strings.Count(host, ":") >= 2 {
		host = "[" + host + "]"
	}

	if err != nil {
		if strings.Contains(err.Error(), "missing port in address") {
			//  according to APISIX upstream nodes policy, port is optional
			host = key
			port = "0"
		} else {
			return nil, errors.New("invalid upstream node")
		}
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Errorf("parse port to int fail: %s", err)
		return nil, err
	}

	node := &Node{
		Host:   host,
		Port:   portInt,
		Weight: int(val),
	}

	return node, nil
}

func NodesFormat(obj interface{}) interface{} {
	nodes := make([]*Node, 0)
	switch objType := obj.(type) {
	case map[string]float64:
		log.Infof("nodes type: %v", objType)
		value := obj.(map[string]float64)
		for key, val := range value {
			node, err := mapKV2Node(key, val)
			if err != nil {
				return obj
			}
			nodes = append(nodes, node)
		}
		return nodes
	case map[string]interface{}:
		log.Infof("nodes type: %v", objType)
		value := obj.(map[string]interface{})
		for key, val := range value {
			node, err := mapKV2Node(key, val.(float64))
			if err != nil {
				return obj
			}
			nodes = append(nodes, node)
		}
		return nodes
	case []*Node:
		log.Infof("nodes type: %v", objType)
		return obj
	case []interface{}:
		log.Infof("nodes type []interface{}: %v", objType)
		list := obj.([]interface{})
		for _, v := range list {
			val := v.(map[string]interface{})
			node := &Node{
				Host:   val["host"].(string),
				Port:   int(val["port"].(float64)),
				Weight: int(val["weight"].(float64)),
			}
			if _, ok := val["priority"]; ok {
				node.Priority = int(val["priority"].(float64))
			}

			if _, ok := val["metadata"]; ok {
				node.Metadata = val["metadata"].(map[string]interface{})
			}

			nodes = append(nodes, node)
		}
		return nodes
	}

	return obj
}
