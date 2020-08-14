package longest_common_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type input struct {
		text1 string
		text2 string
	}
	tests := map[string]struct {
		input input
		want  int
	}{
		"1": {input: input{text1: "abcde", text2: "ace"}, want: 3},
		"2": {input: input{text1: "abc", text2: "abc"}, want: 3},
		"3": {input: input{text1: "abc", text2: "def"}, want: 0},
		"4": {input: input{text1: "psnw", text2: "vozsh"}, want: 1},
		"5": {input: input{text1: "bsbininm", text2: "jmjkbkjkv"}, want: 1},
	}

	/*
		"bsbininm"
		"jmjkbkjkv"
	*/

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, longestCommonSubsequence(tc.input.text1, tc.input.text2))
		})
	}
}
