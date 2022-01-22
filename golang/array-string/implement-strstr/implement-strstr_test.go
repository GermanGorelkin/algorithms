package implement_strstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := map[string]struct {
		input args
		want  int
	}{
		"haystack is empty": {input: args{
			haystack: "",
			needle:   "bba",
		}, want: -1},
		"needle is empty": {input: args{
			haystack: "hello",
			needle:   "",
		}, want: 0},
		"1": {input: args{
			haystack: "hello",
			needle:   "ll",
		}, want: 2},
		"2": {input: args{
			haystack: "aaaaa",
			needle:   "bba",
		}, want: -1},
		"3": {input: args{
			haystack: "aaa",
			needle:   "aaaa",
		}, want: -1},
		"4": {input: args{
			haystack: "mississippi",
			needle:   "sipp",
		}, want: 6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, strStr(tc.input.haystack, tc.input.needle))
		})
	}
}
