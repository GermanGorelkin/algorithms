package find_first_and_last_position_of_element_in_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := map[string]struct {
		args args
		want []int
	}{
		"1": {args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8}, want: []int{3, 4}},
		"2": {args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 6}, want: []int{-1, -1}},
		"3": {args: args{nums: []int{}, target: 0}, want: []int{-1, -1}},
		"4": {args: args{nums: []int{2, 2}, target: 3}, want: []int{-1, -1}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, searchRange(tc.args.nums, tc.args.target))
		})
	}
}
