package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	preIndex := make([]int, 3001)
	inIndex := make([]int, 3001)

	n := len(preorder)
	for i := 0; i < n; i++ {
		preIndex[preorder[i]] = i
		inIndex[inorder[i]] = i
	}

	return nil
}
