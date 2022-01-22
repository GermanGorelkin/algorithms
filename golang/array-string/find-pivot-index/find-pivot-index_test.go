package find_pivot_index

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"[1, 7, 3, 6, 5, 6]": {input: []int{1, 7, 3, 6, 5, 6}, want: 3},
		"[1, 2, 3]":          {input: []int{1, 2, 3}, want: -1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, pivotIndex(tc.input))
		})
	}

}
