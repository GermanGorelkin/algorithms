package search_in_rotated_sorted_array

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
		"1": {input: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		"2": {input: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, want: -1},
		"3": {input: args{nums: []int{3, 1}, target: 1}, want: 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, search(tc.input.nums, tc.input.target))
		})
	}
}

func Test_findRotMid(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{4, 5, 6, 7, 0, 1, 2}, want: 4},
		"2": {input: []int{4, 5, 6, 0, 1, 2}, want: 3},
		"3": {input: []int{0, 1, 2, 3, 4, 5, 6}, want: 0},
		"4": {input: []int{3, 1}, want: 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findRotMid(tc.input))
		})
	}
}
