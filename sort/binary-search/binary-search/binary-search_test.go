package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := map[string]struct {
		input args
		want  int
	}{
		"1": {input: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		"2": {input: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: -1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, search(tc.input.nums, tc.input.target))
		})
	}
}
