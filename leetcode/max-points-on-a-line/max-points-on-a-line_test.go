package maxpointsonaline

import (
	"testing"
)

func Test_maxPoints(t *testing.T) {
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
			name: "all points on a horizontal line",
			args: args{points: [][]int{{1, 1}, {2, 1}, {3, 1}}},
			want: 3,
		},
		{
			name: "all points on a vertical line",
			args: args{points: [][]int{{2, 3}, {2, 4}, {2, 5}}},
			want: 3,
		},
		{
			name: "all points on a diagonal line",
			args: args{points: [][]int{{1, 1}, {2, 2}, {3, 3}}},
			want: 3,
		},
		{
			name: "two points only",
			args: args{points: [][]int{{0, 0}, {1, 1}}},
			want: 2,
		},
		{
			name: "single point",
			args: args{points: [][]int{{1, 1}}},
			want: 1,
		},
		{
			name: "three points, two on a line",
			args: args{points: [][]int{{1, 1}, {2, 2}, {3, 4}}},
			want: 2,
		},
		{
			name: "multiple points, max three on a line",
			args: args{points: [][]int{{1, 1}, {2, 2}, {3, 3}, {3, 4}, {5, 6}}},
			want: 3,
		},
		{
			name: "leetcode example",
			args: args{points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSameDirection(t *testing.T) {
	type args struct {
		a Point
		b Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "same direction positive slope",
			args: args{
				a: Point{x: 2, y: 2},
				b: Point{x: 4, y: 4},
			},
			want: true,
		},
		{
			name: "same direction negative slope",
			args: args{
				a: Point{x: -2, y: -2},
				b: Point{x: -4, y: -4},
			},
			want: true,
		},
		{
			name: "different direction",
			args: args{
				a: Point{x: 1, y: 2},
				b: Point{x: 2, y: 3},
			},
			want: false,
		},
		{
			name: "vertical and horizontal",
			args: args{
				a: Point{x: 0, y: 2},
				b: Point{x: 2, y: 0},
			},
			want: false,
		},
		{
			name: "zero vector and nonzero vector",
			args: args{
				a: Point{x: 0, y: 0},
				b: Point{x: 1, y: 1},
			},
			want: true,
		},
		{
			name: "both zero vectors",
			args: args{
				a: Point{x: 0, y: 0},
				b: Point{x: 0, y: 0},
			},
			want: true,
		},
		{
			name: "test with (-2,-1) and (-3,0)",
			args: args{
				a: Point{x: -2, y: -1},
				b: Point{x: -3, y: 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameDirection(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isSameDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
