package two_sum_ii_input_array_is_sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := map[string]struct {
		input args
		want  []int
	}{
		"1": {input: args{numbers: []int{2, 7, 11, 15}, target: 9}, want: []int{1, 2}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, twoSum(tc.input.numbers, tc.input.target))
		})
	}
}
