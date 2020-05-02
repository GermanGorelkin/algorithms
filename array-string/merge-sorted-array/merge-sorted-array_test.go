package merge_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := map[string]struct {
		input args
		want  []int
	}{
		"1": {input: args{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		}, want: []int{1, 2, 2, 3, 5, 6}},
		"2": {input: args{
			nums1: []int{4, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
		}, want: []int{1, 2, 3, 4, 5, 6}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			merge(tc.input.nums1, tc.input.m, tc.input.nums2, tc.input.n)
			assert.Equal(t, tc.want, tc.input.nums1)
		})
	}
}
