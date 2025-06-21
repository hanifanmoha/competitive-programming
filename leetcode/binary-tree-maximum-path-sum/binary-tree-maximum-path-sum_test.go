package binarytreemaximumpathsum

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "single node",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "simple tree",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 6, // 2 + 1 + 3
		},
		{
			name: "tree with negative values",
			args: args{
				root: &TreeNode{
					Val: -10,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: 42, // 15 + 20 + 7
		},
		{
			name: "all negative values",
			args: args{
				root: &TreeNode{
					Val:   -3,
					Left:  &TreeNode{Val: -2},
					Right: &TreeNode{Val: -1},
				},
			},
			want: -1,
		},
		{
			name: "left skewed tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: 6, // 3 + 2 + 1
		},
		{
			name: "right skewed tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: 6, // 1 + 2 + 3
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
