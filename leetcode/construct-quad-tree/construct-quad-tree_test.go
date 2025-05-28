package constructquadtree

import (
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "single leaf 1x1 grid with 1",
			args: args{grid: [][]int{{1}}},
			want: &Node{Val: true, IsLeaf: true},
		},
		{
			name: "single leaf 1x1 grid with 0",
			args: args{grid: [][]int{{0}}},
			want: &Node{Val: false, IsLeaf: true},
		},
		{
			name: "2x2 grid all 1s",
			args: args{grid: [][]int{
				{1, 1},
				{1, 1},
			}},
			want: &Node{Val: true, IsLeaf: true},
		},
		{
			name: "2x2 grid all 0s",
			args: args{grid: [][]int{
				{0, 0},
				{0, 0},
			}},
			want: &Node{Val: false, IsLeaf: true},
		},
		{
			name: "2x2 grid mixed",
			args: args{grid: [][]int{
				{1, 0},
				{0, 1},
			}},
			want: &Node{
				Val:         false,
				IsLeaf:      false,
				TopLeft:     &Node{Val: true, IsLeaf: true},
				TopRight:    &Node{Val: false, IsLeaf: true},
				BottomLeft:  &Node{Val: false, IsLeaf: true},
				BottomRight: &Node{Val: true, IsLeaf: true},
			},
		},
		{
			name: "8x8 grid mixed values",
			args: args{grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 1},
			}},
			want: &Node{
				Val:    false,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
				BottomLeft: &Node{
					Val:    false,
					IsLeaf: false,
					TopLeft: &Node{
						Val:         false,
						IsLeaf:      false,
						TopLeft:     &Node{Val: true, IsLeaf: true},
						TopRight:    &Node{Val: true, IsLeaf: true},
						BottomLeft:  &Node{Val: true, IsLeaf: true},
						BottomRight: &Node{Val: false, IsLeaf: true},
					},
					TopRight: &Node{
						Val:    false,
						IsLeaf: true,
					},
					BottomLeft: &Node{
						Val:    false,
						IsLeaf: true,
					},
					BottomRight: &Node{
						Val:    false,
						IsLeaf: true,
					},
				},
				BottomRight: &Node{
					Val:         false,
					IsLeaf:      false,
					TopLeft:     &Node{Val: false, IsLeaf: true},
					TopRight:    &Node{Val: false, IsLeaf: true},
					BottomLeft:  &Node{Val: true, IsLeaf: true},
					BottomRight: &Node{Val: true, IsLeaf: true},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}
