package diagonal_traverse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  []int
	}{
		"1": {
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		"2": {
			input: [][]int{
				{2, 5},
				{8, 4},
				{0, -1}},
			want: []int{2, 5, 8, 0, 4, -1}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findDiagonalOrder(tc.input))
		})
	}
}
