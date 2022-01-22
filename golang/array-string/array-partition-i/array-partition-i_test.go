package array_partition_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arrayPairSum(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {
			input: []int{1, 4, 3, 2},
			want:  4},
		"2": {
			input: []int{7, 3, 1, 0, 0, 6},
			want:  7},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, arrayPairSum(tc.input))
		})
	}
}
