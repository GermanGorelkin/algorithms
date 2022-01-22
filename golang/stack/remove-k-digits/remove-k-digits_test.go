package remove_k_digits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := map[string]struct {
		input args
		want  string
	}{
		"1": {input: args{num: "1432219", k: 3}, want: "1219"},
		"2": {input: args{num: "10200", k: 1}, want: "200"},
		"3": {input: args{num: "10", k: 2}, want: "0"},
		"4": {input: args{num: "10", k: 1}, want: "0"},
		"5": {input: args{num: "112", k: 1}, want: "11"},
		"6": {input: args{num: "1234567890", k: 9}, want: "0"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, removeKdigits(tc.input.num, tc.input.k))
		})
	}
}
