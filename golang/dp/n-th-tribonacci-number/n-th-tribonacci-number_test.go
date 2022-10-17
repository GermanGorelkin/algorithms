package n_th_tribonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tribonacci(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "0", input: 0, want: 0},
		{name: "1", input: 1, want: 1},
		{name: "2", input: 2, want: 1},
		{name: "3", input: 3, want: 2},
		{name: "25", input: 25, want: 1389537},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tribonacci(tc.input))
		})
	}
}
