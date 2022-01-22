package insertion


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}

		sortArray(arr)

		assert.Equal(t, expected, arr)
	})
	t.Run("2", func(t *testing.T) {
		arr := []int{5}
		expected := []int{5}

		sortArray(arr)

		assert.Equal(t, expected, arr)
	})
	t.Run("3", func(t *testing.T) {
		arr := []int{5, 1}
		expected := []int{1, 5}

		sortArray(arr)

		assert.Equal(t, expected, arr)
	})
	t.Run("4", func(t *testing.T) {
		arr := []int{1, 5}
		expected := []int{1, 5}

		sortArray(arr)

		assert.Equal(t, expected, arr)
	})
}


