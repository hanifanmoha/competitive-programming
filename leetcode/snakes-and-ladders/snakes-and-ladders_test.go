package snakes_and_ladders

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{
			name: "Test case 1",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 35, -1, -1, 13, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 15, -1, -1, -1, -1},
				},
			},
			want: 4,
		},
		{
			name: "Test case 2",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 35, -1, -1, 13, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 15, -1, -1, -1, 2},
				},
			},
			want: 3,
		},
		{
			name: "Edge case 1: Empty board",
			args: args{
				board: [][]int{},
			},
			want: -1,
		},
		{
			name: "Edge case 2: Single cell board",
			args: args{
				board: [][]int{
					{-1},
				},
			},
			want: 0,
		},
		{
			name: "Edge case 3: No ladders or snakes",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.args.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
