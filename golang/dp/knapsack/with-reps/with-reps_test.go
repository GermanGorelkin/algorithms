package with_reps

import (
	"testing"
)

func Test_knapsack(t *testing.T) {
	type args struct {
		W int
		w []int
		c []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				W: 10,
				w: []int{6, 3, 4, 2},
				c: []int{30, 14, 16, 9},
			},
			want: 48,
		},
		{
			name: "2",
			args: args{
				W: 10,
				w: []int{5, 3},
				c: []int{3, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.W, tt.args.w, tt.args.c); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
