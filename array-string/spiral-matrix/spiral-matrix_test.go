package spiral_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  []int
	}{
		"1": {input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		"2": {input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, spiralOrder(tc.input))
		})
	}
}
