package reverse_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits(t *testing.T) {
	tests := map[string]struct {
		args uint32
		want uint32
	}{
		"43261596": {args: 43261596, want: 964176192},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, reverseBits(tc.args))
		})
	}
}