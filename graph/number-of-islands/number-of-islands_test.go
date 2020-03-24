package number_of_islands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		grid := [][]byte{
			{1, 1, 1, 1, 0},
			{1, 1, 0, 1, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		want := 1
		assert.Equal(t, want, numIslands(grid))
	})
	t.Run("2", func(t *testing.T) {
		grid := [][]byte{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		}
		want := 3
		assert.Equal(t, want, numIslands(grid))
	})
	t.Run("3", func(t *testing.T) {
		var grid [][]byte
		want := 0
		assert.Equal(t, want, numIslands(grid))
	})
	t.Run("4", func(t *testing.T) {
		grid := [][]byte{
			{1, 0, 1, 0, 1},
		}
		want := 3
		assert.Equal(t, want, numIslands(grid))
	})
	t.Run("5", func(t *testing.T) {
		grid := [][]byte{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		}
		want := 1
		assert.Equal(t, want, numIslands(grid))
	})
}
