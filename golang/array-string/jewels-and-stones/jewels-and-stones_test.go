package jewels_and_stones

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := map[string]struct {
		input args
		want  int
	}{
		"1": {input: args{J: "aA", S: "aAAbbbb"}, want: 3},
		"2": {input: args{J: "z", S: "ZZ"}, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, numJewelsInStones(tc.input.J, tc.input.S), tc.want)
		})
	}
}
