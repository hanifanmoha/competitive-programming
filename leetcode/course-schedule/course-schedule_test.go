package courseschedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Example 1 - possible to finish",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "Example 2 - impossible to finish (cycle)",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "No prerequisites",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{},
			},
			want: true,
		},
		{
			name: "Multiple independent chains",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			},
			want: true,
		},
		{
			name: "Cycle in longer chain",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			},
			want: false,
		},
		{
			name: "Disconnected graph, one cycle",
			args: args{
				numCourses:    5,
				prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 4}},
			},
			want: false,
		},
		{
			name: "Disconnected graph, no cycles",
			args: args{
				numCourses:    5,
				prerequisites: [][]int{{1, 0}, {3, 4}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
