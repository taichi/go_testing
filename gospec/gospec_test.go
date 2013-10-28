package go_testing_test

import (
	. ".."
	. "github.com/orfjackal/gospec/src/gospec"
)

func newmathSpec(c Context) {
	c.Specify("The value should sqrt number", func() {
		c.Expect(Sqrt(4), Equals, float64(2))
	})
	c.Specify("The value should fail assertion", func() {
		c.Expect(Sqrt(4), Equals, 2)
	})
}
