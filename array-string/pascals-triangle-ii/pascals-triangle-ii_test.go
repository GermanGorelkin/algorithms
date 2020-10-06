package pascals_triangle_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getRow(t *testing.T) {
	want := []int{1, 3, 3, 1}

	assert.Equal(t, want, getRow(3))
}
