package go_testing_test

import (
	. ".."
	"github.com/bmizerany/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	assert.Equal(t, Sqrt(4), float64(2), "they should be equal")
	assert.Equal(t, Sqrt(4), 2, "they should not be equal")
}
