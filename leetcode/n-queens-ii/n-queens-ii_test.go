package nqueensii

import (
	"testing"
)

func Test_isSameDiagonal(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{
			name: "Same diagonal positive slope",
			args: args{x1: 0, y1: 0, x2: 1, y2: 1},
			want: true,
		},
		{
			name: "Same diagonal negative slope",
			args: args{x1: 2, y1: 0, x2: 0, y2: 2},
			want: true,
		},
		{
			name: "Not same diagonal",
			args: args{x1: 0, y1: 0, x2: 1, y2: 2},
			want: false,
		},
		{
			name: "Same position",
			args: args{x1: 3, y1: 3, x2: 3, y2: 3},
			want: true,
		},
		{
			name: "Far apart same diagonal",
			args: args{x1: 0, y1: 7, x2: 7, y2: 0},
			want: true,
		},
		{
			name: "Far apart not same diagonal",
			args: args{x1: 0, y1: 7, x2: 7, y2: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameDiagonal(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("isSameDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalNQueens(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{
			name: "N=1",
			args: args{N: 1},
			want: 1,
		},
		{
			name: "N=2",
			args: args{N: 2},
			want: 0,
		},
		{
			name: "N=3",
			args: args{N: 3},
			want: 0,
		},
		{
			name: "N=4",
			args: args{N: 4},
			want: 2,
		},
		{
			name: "N=5",
			args: args{N: 5},
			want: 10,
		},
		{
			name: "N=6",
			args: args{N: 6},
			want: 4,
		},
		{
			name: "N=7",
			args: args{N: 7},
			want: 40,
		},
		{
			name: "N=8",
			args: args{N: 8},
			want: 92,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.N); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
