package go_testing_test

import (
	. ".."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	Describe("newmath", func() {
		Context("Sqrt", func() {
			It("The value should sqrt number", func() {
				Expect(Sqrt(4)).To(Equal(float64(2)))
			})
			It("The value should fail assertion", func() {
				Ω(Sqrt(4)).To(Equal(2))
			})
		})
	})
}
