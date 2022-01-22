package v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  []byte
	}{
		"1": {input: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reverseString(tc.input)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}
