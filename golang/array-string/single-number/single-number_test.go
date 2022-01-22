package single_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"[2,2,1]":     {input: []int{2, 2, 1}, want: 1},
		"[4,1,2,1,2]": {input: []int{4, 1, 2, 1, 2}, want: 4},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, singleNumber(tc.input))
		})
	}
}
