package flood_fill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}

	tests := map[string]struct {
		input args
		want  [][]int
	}{
		"1": {
			input: args{
				image: [][]int{
					{1, 1, 1}, {1, 1, 0}, {1, 0, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 2},
			want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, floodFill(tc.input.image, tc.input.sr, tc.input.sc, tc.input.newColor))
		})
	}
}
