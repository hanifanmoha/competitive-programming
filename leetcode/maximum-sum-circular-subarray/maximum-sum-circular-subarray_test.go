package maximumsumcircularsubarray

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "10 elements, mixed positive and negative",
			args: args{nums: []int{8, -1, 3, 4, -5, 6, -2, 7, -3, 2}},
			want: 24,
		},
		{
			name: "10 elements, all positive",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "10 elements, all negative",
			args: args{nums: []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}},
			want: -1,
		},
		{
			name: "10 elements, alternating sign",
			args: args{nums: []int{5, -2, 5, -2, 5, -2, 5, -2, 5, -2}},
			want: 17,
		},
		{
			name: "10 elements, zeros and negatives",
			args: args{nums: []int{0, -1, 0, -2, 0, -3, 0, -4, 0, -5}},
			want: 0,
		},
		{
			name: "all positive numbers",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "all negative numbers",
			args: args{nums: []int{-3, -2, -1, -4}},
			want: -1,
		},
		{
			name: "mixed numbers, max subarray not circular",
			args: args{nums: []int{5, -3, 5}},
			want: 10,
		},
		{
			name: "mixed numbers, max subarray is circular",
			args: args{nums: []int{3, -1, 2, -1}},
			want: 4,
		},
		{
			name: "single element",
			args: args{nums: []int{7}},
			want: 7,
		},
		{
			name: "wrap around needed",
			args: args{nums: []int{2, -2, 2, 7, 8, 0}},
			want: 19,
		},
		{
			name: "all zeros",
			args: args{nums: []int{0, 0, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
