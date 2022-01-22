package minimum_path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		grid := [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}

		assert.Equal(t, 7, minPathSum(grid))
	})
}
