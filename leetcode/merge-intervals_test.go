package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "no intervals",
			args: args{intervals: [][]int{}},
			want: [][]int{},
		},
		{
			name: "single interval",
			args: args{intervals: [][]int{{1, 3}}},
			want: [][]int{{1, 3}},
		},
		{
			name: "unsorted intervals",
			args: args{intervals: [][]int{{5, 6}, {1, 3}, {8, 10}, {2, 4}, {15, 18}, {7, 9}, {12, 14}, {11, 13}, {3, 5}, {6, 7}}},
			want: [][]int{{1, 10}, {11, 14}, {15, 18}},
		},
		{
			name: "non-overlapping intervals",
			args: args{intervals: [][]int{{1, 2}, {3, 4}}},
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "overlapping intervals",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "contained intervals",
			args: args{intervals: [][]int{{1, 4}, {2, 3}}},
			want: [][]int{{1, 4}},
		},
		{
			name: "multiple merges",
			args: args{intervals: [][]int{{1, 4}, {0, 2}, {3, 5}}},
			want: [][]int{{0, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
