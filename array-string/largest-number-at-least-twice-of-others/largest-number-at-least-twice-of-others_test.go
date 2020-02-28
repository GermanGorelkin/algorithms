package largest_number_at_least_twice_of_others

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dominantIndex(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"[3, 6, 1, 0]": {input: []int{3, 6, 1, 0}, want: 1},
		"[1, 2, 3, 4]": {input: []int{1, 2, 3, 4}, want: -1},
		"[0, 0, 3, 2]": {input: []int{0, 0, 3, 2}, want: -1},
		"[1]":          {input: []int{1}, want: 0},
		"[1, 0]":       {input: []int{1, 0}, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, dominantIndex(tc.input))
		})
	}
}
