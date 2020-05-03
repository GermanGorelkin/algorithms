package ransom_note

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := map[string]struct {
		input args
		want  bool
	}{
		"1": {input: args{ransomNote: "a", magazine: "b"}, want: false},
		"2": {input: args{ransomNote: "aa", magazine: "ab"}, want: false},
		"3": {input: args{ransomNote: "aa", magazine: "aab"}, want: true},
		"4": {input: args{ransomNote: "", magazine: "abb"}, want: true},
		"5": {input: args{ransomNote: "aa", magazine: ""}, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, canConstruct(tc.input.ransomNote, tc.input.magazine), tc.want)
		})
	}
}
