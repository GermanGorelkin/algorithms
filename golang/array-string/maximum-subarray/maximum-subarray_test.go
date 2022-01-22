package maximum_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArrayV1(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{1}, want: 1},
		"2": {input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		"3": {input: []int{-2, 1}, want: 1},
		"4": {input: []int{-2, -1}, want: -1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxSubArrayV1(tc.input))
		})
	}
}

func Test_maxSubArrayV2(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{1}, want: 1},
		"2": {input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		"3": {input: []int{-2, 1}, want: 1},
		"4": {input: []int{-2, -1}, want: -1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxSubArrayV2(tc.input))
		})
	}
}
