package minimum_time_visiting_all_points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimum_time_visiting_all_points(t *testing.T){
	tests := map[string]struct{
		input [][]int
		want int
	}{
		"1":{
			input: [][]int{{1,1},{3,4},{-1,0}},
			want: 7,
		},
		"2":{
			input: [][]int{{3,2},{-2,2}},
			want: 5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func (t *testing.T){
			assert.Equal(t, tc.want, minTimeToVisitAllPoints(tc.input))
		})
	}
}
