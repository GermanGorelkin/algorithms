package longest_increasing_subseq

import (
	"reflect"
	"testing"
)

func TestLIS1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "2",
			args: args{arr: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "3",
			args: args{arr: []int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS1(tt.args.arr); got != tt.want {
				t.Errorf("LIS1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLIS2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: []int{2, 3, 7, 101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LIS2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLIS3(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: []int{2, 3, 7, 101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS3(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LIS3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLIS4(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "2",
			args: args{arr: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "3",
			args: args{arr: []int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
		{
			name: "4",
			args: args{arr: []int{1, 2, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS4(tt.args.arr); got != tt.want {
				t.Errorf("LIS4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLIS5(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 1}},
			want: 4,
		},
		{
			name: "2",
			args: args{arr: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "3",
			args: args{arr: []int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
		{
			name: "4",
			args: args{arr: []int{1, 2, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS5(tt.args.arr); got != tt.want {
				t.Errorf("LIS5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upper_bound(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{arr: []int{1, 2, 3}, x: 3},
			want: 2,
		},
		{
			name: "2",
			args: args{arr: []int{1, 2, 3}, x: 2},
			want: 2,
		},
		{
			name: "3",
			args: args{arr: []int{1, 2, 2, 3}, x: 2},
			want: 3,
		},
		{
			name: "4",
			args: args{arr: []int{1, 2, 2, 3, 3, 4, 5}, x: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upper_bound(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("upper_bound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLIS6(t *testing.T) {
	t.Skip()
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 1, 18}},
			want: []int{2, 3, 7, 18},
		},
		// {
		// 	name: "1",
		// 	args: args{arr: []int{3, 1, 5, 2, 6, 4, 9}},
		// 	want: []int{1, 2, 4, 9},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS6(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LIS6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lower_bound(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{arr: []int{3, 2, 1}, x: 3},
			want: 1,
		},
		{
			name: "1",
			args: args{arr: []int{5, 4, 3, 2, 1}, x: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lower_bound(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("lower_bound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDS1(t *testing.T) {
	t.Skip()
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: []int{10, 9, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LDS1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LDS1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLIS7(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 1, 18}},
			want: []int{2, 3, 7, 18},
		},
		{
			name: "1",
			args: args{nums: []int{3, 1, 5, 2, 6, 4, 9}},
			want: []int{1, 2, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS7(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LIS7() = %v, want %v", got, tt.want)
			}
		})
	}
}
