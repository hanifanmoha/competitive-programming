package insertinterval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "Test 2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "Test 3",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{5, 7},
			},
			want: [][]int{{5, 7}},
		},
		{
			name: "Test 4",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{2, 3},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "Test 5",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{2, 7},
			},
			want: [][]int{{1, 7}},
		},
		{
			name: "Test 6",
			args: args{
				intervals:   [][]int{{1, 5}, {6, 8}},
				newInterval: []int{0, 9},
			},
			want: [][]int{{0, 9}},
		},
		{
			name: "Test 7",
			args: args{
				intervals:   [][]int{{1, 5}, {6, 8}},
				newInterval: []int{5, 6},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "Test 8",
			args: args{
				intervals:   [][]int{{1, 3}, {4, 6}, {8, 10}},
				newInterval: []int{2, 7},
			},
			want: [][]int{{1, 7}, {8, 10}},
		},
		{
			name: "Test 9",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 6},
			},
			want: [][]int{{1, 9}},
		},
		{
			name: "Test 10",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{17, 19},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {17, 19}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
