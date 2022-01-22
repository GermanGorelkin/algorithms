package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	type input struct {
		n, m int
	}

	tests := map[string]struct {
		input input
		want  int
	}{
		"1": {input: input{n: 3, m: 2}, want: 3},
		"2": {input: input{n: 7, m: 3}, want: 28},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, uniquePaths(tc.input.m, tc.input.n))
		})
	}
}
