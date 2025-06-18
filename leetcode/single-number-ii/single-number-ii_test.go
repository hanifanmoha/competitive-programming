package singlenumberii

import "testing"

func Test_singleNumber(t *testing.T) {
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
			name: "example1",
			args: args{nums: []int{2, 2, 3, 2}},
			want: 3,
		},
		{
			name: "example2",
			args: args{nums: []int{0, 1, 0, 1, 0, 1, 99}},
			want: 99,
		},
		{
			name: "negative numbers",
			args: args{nums: []int{-2, -2, -2, -7}},
			want: -7,
		},
		{
			name: "single element",
			args: args{nums: []int{42}},
			want: 42,
		},
		{
			name: "all same except one",
			args: args{nums: []int{5, 5, 5, 8}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
