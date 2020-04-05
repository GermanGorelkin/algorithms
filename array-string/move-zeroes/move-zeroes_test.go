package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"1": {input: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		"2": {input: []int{0, 0, 1}, want: []int{1, 0, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			moveZeroes(tc.input)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}

func Test_moveZeroesV2(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"1": {input: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		"2": {input: []int{0, 0, 1}, want: []int{1, 0, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			moveZeroesV2(tc.input)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}

func Test_moveZeroesV3(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"1": {input: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		"2": {input: []int{0, 0, 1}, want: []int{1, 0, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			moveZeroesV3(tc.input)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}
