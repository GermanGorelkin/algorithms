package group_anagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  [][]string
	}{
		"1": {input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}},
		"2": {input: []string{"tea", "", "eat", "", "tea", ""}, want: [][]string{
			{"tea", "eat", "tea"},
			{"", "", ""},
		}},
		"3": {input: []string{"tea", "and", "ace", "ad", "eat", "dans"}, want: [][]string{
			{"eat", "tea"},
			{"and"},
			{"dans"},
			{"ace"},
			{"ad"},
		}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, groupAnagrams(tc.input))
		})
	}
}
