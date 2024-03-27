package label_test

import (
	"testing"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/gateway/meta_api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLabel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Label Suite")
}

var _ = BeforeSuite(func() {
	base.CleanAllResource()
	time.Sleep(base.SleepTime)
})
