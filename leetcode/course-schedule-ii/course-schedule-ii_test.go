package coursescheduleii

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "LeetCode Example 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "LeetCode Example 2",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			want: []int{0, 1, 2, 3}, // or [0,2,1,3] is also valid
		},
		{
			name: "LeetCode Example 3 - Unreachable (cycle)",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
		{
			name: "LeetCode Example 4 - Cycle (final block unreachable)",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{0, 1}, {1, 0}},
			},
			want: []int{},
		},
		{
			name: "LeetCode Example 5",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 2},
					{1, 3},
					{3, 2},
					{2, 0},
				},
			},
			want: []int{0, 2, 3, 1},
		},
		{
			name: "Random order with 8 courses, custom prerequisites",
			args: args{
				numCourses: 8,
				prerequisites: [][]int{
					{1, 0},
					{2, 6},
					{1, 7},
					{6, 4},
					{7, 0},
					{0, 5},
				},
			},
			want: []int{5, 4, 6, 3, 2, 0, 7, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
