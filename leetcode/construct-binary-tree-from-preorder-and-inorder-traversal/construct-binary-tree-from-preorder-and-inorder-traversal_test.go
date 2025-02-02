package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "example 2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: &TreeNode{
				Val: -1,
			},
		},
		{
			name: "example 3",
			args: args{
				preorder: []int{1, 2, 4, 5, 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
				inorder:  []int{4, 2, 5, 1, 6, 3, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 8,
							Right: &TreeNode{
								Val: 9,
								Right: &TreeNode{
									Val: 10,
									Right: &TreeNode{
										Val: 11,
										Right: &TreeNode{
											Val: 12,
											Right: &TreeNode{
												Val: 13,
												Right: &TreeNode{
													Val: 14,
													Right: &TreeNode{
														Val: 15,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "example 4",
			args: args{
				preorder: []int{8, 4, 2, 1, 3, 6, 5, 7, 12, 10, 9, 11, 14, 13, 15},
				inorder:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			},
			want: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 12,
					Left: &TreeNode{
						Val: 10,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 11,
						},
					},
					Right: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 15,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
