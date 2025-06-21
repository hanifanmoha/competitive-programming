package convertsortedarraytobinarysearchtree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}

	var p, c *TreeNode = nil, root
	isLeft := false

	for c != nil {
		p = c
		if val < c.Val {
			isLeft = true
			c = c.Left
		} else {
			isLeft = false
			c = c.Right
		}
	}

	if isLeft {
		p.Left = NewTreeNode(val)
	} else {
		p.Right = NewTreeNode(val)
	}

	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	fmt.Println("=>", root.Val)
	traverse(root.Right)
}

/*
1 2 3 4 5 6 7 8 9 10

0 9 -> 4
0 3 -> 1
0 0 -> 0
2 3 -> 2
2 1 -> return
3 3 -> 3
5 9 ->


*/

func binaryInsert(root *TreeNode, nums []int, left, right int) *TreeNode {
	if left > right {
		return root
	}

	mid := left + (right-left)/2
	// fmt.Println(mid)

	root = insert(root, nums[mid])

	binaryInsert(root, nums, left, mid-1)
	binaryInsert(root, nums, mid+1, right)

	return root

}

func sortedArrayToBST(nums []int) *TreeNode {

	// var root *TreeNode

	// for _, x := range nums {
	// 	root = insert(root, x)
	// }

	root := binaryInsert(nil, nums, 0, len(nums)-1)

	traverse(root)

	return root

}
