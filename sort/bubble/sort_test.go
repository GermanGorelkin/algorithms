package sort

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
	"testing/quick"
)

func TestSort(t *testing.T) {
	t.Run("simply", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}

		Sort(arr)

		assert.Equal(t, expected, arr)
	})

	t.Run("random arr", func(t *testing.T) {
		f := func(arr []int) bool {
			expected := make([]int, len(arr))
			copy(expected, arr)
			sort.Ints(expected)

			Sort(arr)

			return reflect.DeepEqual(arr, expected)
		}

		if err := quick.Check(f, nil); err != nil {
			t.Error(err)
		}
	})
}

func TestSort2(t *testing.T) {
	t.Run("simply", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}

		Sort2(arr)

		assert.Equal(t, expected, arr)
	})

	t.Run("random arr", func(t *testing.T) {
		f := func(arr []int) bool {
			expected := make([]int, len(arr))
			copy(expected, arr)
			sort.Ints(expected)

			Sort2(arr)

			return reflect.DeepEqual(arr, expected)
		}

		if err := quick.Check(f, nil); err != nil {
			t.Error(err)
		}
	})
}
