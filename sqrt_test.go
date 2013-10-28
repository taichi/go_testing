package go_testing_test

import (
	. "."
	"testing"
)

func doTest(t *testing.T, in float64, out float64) {
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestSqrt(t *testing.T) {
	doTest(t, 4, 2)
	doTest(t, 4, 22)
}
