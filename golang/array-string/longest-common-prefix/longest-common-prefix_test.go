package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"1": {input: []string{"flower", "flow", "flight"}, want: "fl"},
		"2": {input: []string{"dog", "racecar", "car"}, want: ""},
		"3": {input: []string{"c", "c"}, want: "c"},
		"4": {input: []string{"aa", "a"}, want: "a"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, longestCommonPrefix(tc.input))
		})
	}
}
