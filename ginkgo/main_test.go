package ginkgo

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Describe", func() {
	BeforeEach(func() {
		fmt.Println("BeforeEach in Describe")
	})

	Context("Hello", func() {
		BeforeEach(func() {
			fmt.Println("BeforeEach in Context")
		})
		It("It", func() {
			defer func() {
				fmt.Println("hello")
			}()
			somefunc()
			Expect(false).To(BeTrue())
		})

	})

	When("World", func() {
		BeforeEach(func() {
			fmt.Println("BeforeEach in When")
			DeferCleanup(func() {
				fmt.Println("DeferCleanup in When")
			})
		})
		It("It2", func() {
			fmt.Println("world")
		})
		AfterEach(func() {
			fmt.Println("AfterEach in When")
		})
	})
	AfterEach(func() {
		fmt.Println("AfterEach in Describe")
	})
})

func somefunc() {
	str := "helldsfgs dvfd"
	Expect(false).To(BeTrue(), "in somefunc %s", str)

}
