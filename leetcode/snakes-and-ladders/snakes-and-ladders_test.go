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
			want: 4,
		},
		{
			name: "Test case 3",
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
			want: 6,
		},
		{
			name: "Test case 4",
			args: args{
				board: [][]int{
					{-1, 83, -1, 46, -1, -1, -1, -1, 40, -1},
					{-1, 29, -1, -1, -1, 51, -1, 18, -1, -1},
					{-1, 35, 31, 51, -1, 6, -1, 40, -1, -1},
					{-1, -1, -1, 28, -1, 36, -1, -1, -1, -1},
					{-1, -1, -1, -1, 44, -1, -1, 84, -1, -1},
					{-1, -1, -1, 31, -1, 98, 27, 94, 74, -1},
					{4, -1, -1, 46, 3, 14, 7, -1, 84, 67},
					{-1, -1, -1, -1, 2, 72, -1, -1, 86, -1},
					{-1, 32, -1, -1, -1, -1, -1, -1, -1, 19},
					{-1, -1, -1, -1, -1, 72, 46, -1, 92, 6},
				},
			},
			want: 3,
		},
		{
			name: "Test case 5",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{6, 5, 4, 3, 2, 1},
					{-1, -1, -1, -1, -1, -1},
				},
			},
			want: -1,
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
