package check_if_it_is_a_straight_line

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkStraightLine(t *testing.T) {
	tests := map[string]struct {
		input [][]int
		want  bool
	}{
		"1": {input: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, want: true},
		"2": {input: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}, want: true},
		"3": {input: [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, want: false},
		"4": {input: [][]int{{0, 0}, {5, 10}}, want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, checkStraightLine(tc.input))
		})
	}
}
