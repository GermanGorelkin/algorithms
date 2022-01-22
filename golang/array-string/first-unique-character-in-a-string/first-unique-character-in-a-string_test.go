package first_unique_character_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"1": {input: "leetcode", want: 0},
		"2": {input: "loveleetcode", want: 2},
		"3": {input: "lbgtddtgbl", want: -1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, firstUniqChar(tc.input))
		})
	}
}
