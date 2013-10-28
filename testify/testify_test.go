package go_testing_test

import (
	. ".."
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestify(t *testing.T) {
	assert.Equal(t, Sqrt(4), 2, "they should be equal")
	assert.Equal(t, Sqrt(4), 22, "they should not be equal")
}
