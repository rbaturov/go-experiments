package ginkgo

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = AfterSuite(func() {
	fmt.Printf("after suite")
})

func Test5LatencyTesting(t *testing.T) {
	defer teardown()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test5LatencyTesting")
}

func teardown() {
	defer GinkgoRecover()
	fmt.Printf("teardown")
	Expect(false).To(BeTrue(), "in teardown")
}
