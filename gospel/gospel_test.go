package go_testing_test

import (
	. ".."
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T) {
	Describe(t, "newmath", func() {
		Context("Sqrt", func() {
			It("The value should sqrt number", func() {
				Expect(Sqrt(4)).To(Equal, float64(2))
			})
			It("The value should fail assertion", func() {
				Expect(Sqrt(4)).To(Equal, 22)
			})
		})
	})
}
