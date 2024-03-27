package data_loader_test

import (
	"testing"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDataLoader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Loader Suite")
}

var _ = BeforeSuite(func() {
	base.CleanAllResource()
	base.RestartManagerAPI()
	time.Sleep(base.SleepTime)
})
