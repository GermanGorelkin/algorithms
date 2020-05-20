package min_cost_climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{10, 15, 20}, want: 15},
		"2": {input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, want: 6},
		"3": {input: []int{1, 100}, want: 1},
		"4": {input: []int{100, 1}, want: 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, minCostClimbingStairs(tc.input))
		})
	}
}
