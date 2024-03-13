package router

import (
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

var resourceList = make([]*pb.Resource, 0)

func LoadResourceList(resources []*pb.Resource) {
	resourceList = resources
}
