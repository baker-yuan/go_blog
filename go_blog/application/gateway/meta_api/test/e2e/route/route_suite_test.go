package route_test

import (
	"testing"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Route Suite")
}

var _ = BeforeSuite(func() {
	base.CleanAllResource()
	base.RestartManagerAPI()
	time.Sleep(5 * time.Second)
})
