package happy_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isHappy(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"19": {input: 19, want: true},
		"11": {input: 11, want: false},
		"7":  {input: 7, want: true},
		"1":  {input: 1, want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, isHappy(tc.input))
		})
	}
}

func Test_split(t *testing.T) {
	tests := map[string]struct {
		input int
		want  []int
	}{
		"1":    {input: 1, want: []int{1}},
		"19":   {input: 19, want: []int{9, 1}},
		"123":  {input: 123, want: []int{3, 2, 1}},
		"1234": {input: 1234, want: []int{4, 3, 2, 1}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, split(tc.input))
		})
	}
}
