package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	Hi()
	// assert.Equal(t, 123, 125, "they should be equal")
	assert.NotEqual(t, 123, 123, "they should not be equal")
}
