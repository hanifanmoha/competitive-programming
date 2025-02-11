package minimum_number_of_arrows_to_burst_balloons

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test_case_1",
			args: args{points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}},
			want: 2,
		},
		{
			name: "test_case_2",
			args: args{points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
			want: 4,
		},
		{
			name: "test_case_3",
			args: args{points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}},
			want: 2,
		},
		{
			name: "test_case_4",
			args: args{points: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}}},
			want: 3,
		},
		{
			name: "test_case_5",
			args: args{points: [][]int{{1, 2}}},
			want: 1,
		},
		{
			name: "test_case_6",
			args: args{points: [][]int{{2, 3}, {2, 3}}},
			want: 1,
		},
		{
			name: "test_case_7",
			args: args{points: [][]int{{1, 2}, {2, 3}, {3, 4}}},
			want: 2,
		},
		{
			name: "test_case_8",
			args: args{points: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}}},
			want: 1,
		},
		{
			name: "test_case_9",
			args: args{points: [][]int{{1, 2}, {3, 4}, {5, 6}}},
			want: 3,
		},
		{
			name: "test_case_10",
			args: args{points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
