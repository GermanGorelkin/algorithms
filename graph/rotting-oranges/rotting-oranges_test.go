package rotting_oranges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		grid := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}

		assert.Equal(t, 4, orangesRotting(grid))
	})
	t.Run("2", func(t *testing.T) {
		grid := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		}

		assert.Equal(t, -1, orangesRotting(grid))
	})
	t.Run("3", func(t *testing.T) {
		grid := [][]int{
			{0, 2},
		}

		assert.Equal(t, 0, orangesRotting(grid))
	})
	t.Run("4", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 2, 0, 2, 0},
		}

		assert.Equal(t, 2, orangesRotting(grid))
	})
	t.Run("5", func(t *testing.T) {
		grid := [][]int{
			{1, 2, 1, 1, 2, 1, 1},
		}

		assert.Equal(t, 2, orangesRotting(grid))
	})
	t.Run("5", func(t *testing.T) {
		grid := [][]int{
			{2, 2, 2, 1, 1},
		}

		assert.Equal(t, 2, orangesRotting(grid))
	})
	t.Run("5", func(t *testing.T) {
		grid := [][]int{
			{2, 2},
			{1, 1},
			{0, 0},
			{2, 0},
		}
		assert.Equal(t, 1, orangesRotting(grid))
	})
}
