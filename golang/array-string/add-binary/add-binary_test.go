package add_binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := map[string]struct {
		input args
		want  string
	}{
		"1": {
			input: args{a: "11", b: "1"},
			want:  "100",
		},
		"2": {
			input: args{a: "1010", b: "1011"},
			want:  "10101",
		},
		"3": {
			input: args{a: "100", b: "110010"},
			want:  "110110",
		},
		"4": {
			input: args{a: "110010", b: "10111"},
			want:  "1001001",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, addBinary(tc.input.a, tc.input.b))
		})
	}
}

//[49 49 0 49 49 48]
//[49 49 48 49 49 48]
