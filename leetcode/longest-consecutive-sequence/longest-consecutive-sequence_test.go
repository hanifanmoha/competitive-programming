package longest_consecutive_sequence

import "testing"

func Test_longestConsecutive(t *testing.T) {
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
			name: "example 1",
			args: args{nums: []int{100, 4, 200, 1, 3, 2}},
			want: 4,
		},
		{
			name: "example 2",
			args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			want: 9,
		},
		{
			name: "empty array",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "no consecutive sequence",
			args: args{nums: []int{10, 5, 100}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
