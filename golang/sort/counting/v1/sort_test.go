package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("simply", func(t *testing.T) {
		arr := []int{5, 5, 4, 3, 3, 2, 1}
		expected := []int{1, 2, 3, 3, 4, 5, 5}

		Sort(arr)

		assert.Equal(t, expected, arr)
	})
}
