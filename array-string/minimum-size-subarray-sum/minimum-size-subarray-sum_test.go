package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := map[string]struct {
		input args
		want  int
	}{
		"1": {input: args{s: 7, nums: []int{2, 3, 1, 2, 4, 3}}, want: 2},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, minSubArrayLen(tc.input.s, tc.input.nums))
		})
	}
}
