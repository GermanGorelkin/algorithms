package find_numbers_with_even_number_of_digits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"[12,345,2,6,7896]":  {input: []int{12, 345, 2, 6, 7896}, want: 2},
		"[555,901,482,1771]": {input: []int{555, 901, 482, 1771}, want: 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findNumbers(tc.input))
		})
	}
}
