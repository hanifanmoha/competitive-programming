package convertsortedarraytobinarysearchtree

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "empty array",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: &TreeNode{Val: 1},
		},
		{
			name: "two elements",
			args: args{nums: []int{1, 2}},
			want: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			name: "three elements",
			args: args{nums: []int{1, 2, 3}},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "four elements",
			args: args{nums: []int{1, 2, 3, 4}},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 4},
			},
		},
		{
			name: "five elements",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   4,
					Right: &TreeNode{Val: 5},
				},
			},
		},
		{
			name: "twenty elements",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
			want: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 1},
						},
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 5},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val:   9,
							Right: &TreeNode{Val: 10},
						},
					},
				},
				Right: &TreeNode{
					Val: 16,
					Left: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val:   12,
							Right: &TreeNode{Val: 13},
						},
						Right: &TreeNode{
							Val: 15,
						},
					},
					Right: &TreeNode{
						Val: 18,
						Left: &TreeNode{
							Val: 17,
						},
						Right: &TreeNode{
							Val:   19,
							Right: &TreeNode{Val: 20},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
