package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := map[string]struct {
		args args
		want [][]int
	}{
		"1": {
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, permute(tc.args.nums))
		})
	}
}
