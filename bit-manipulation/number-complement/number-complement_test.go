package number_complement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findComplement(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"2": {input: 5, want: 2},
		"1": {input: 1, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findComplement(tc.input))
		})
	}
}
