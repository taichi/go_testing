package go_testing_test

import (
	. ".."
	. "github.com/franela/goblin"
	"testing"
)

func TestGoblin(t *testing.T) {
	g := Goblin(t)
	g.Describe("Goblin", func() {
		g.It("Should sqrt number ", func() {
			g.Assert(Sqrt(4)).Equal(2)
		})
		g.It("Should fail ", func() {
			g.Assert(Sqrt(4)).Equal(22)
		})
	})
}
