package searchinrotatedsortedarray

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{1, 3},
				target: 1,
			},
			want: 0,
		},
		{
			name: "test5",
			args: args{
				nums:   []int{1, 3},
				target: 3,
			},
			want: 1,
		},
		{
			name: "test6",
			args: args{
				nums:   []int{5, 1, 3},
				target: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchStart(t *testing.T) {
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
			name: "test1",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				nums: []int{6, 7, 1, 2, 3, 4, 5},
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				nums: []int{2, 3, 4, 5, 6, 7, 1},
			},
			want: 6,
		},
		{
			name: "test5",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "test5",
			args: args{
				nums: []int{5, 1, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchStart(tt.args.nums); got != tt.want {
				t.Errorf("searchStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
