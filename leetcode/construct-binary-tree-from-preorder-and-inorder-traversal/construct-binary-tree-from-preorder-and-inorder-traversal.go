package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	node := &TreeNode{}
	node.Val = val
	return node
}

func inOrder(arr *[]int, root *TreeNode, isLeft bool, lvl int) {
	if root == nil {
		return
	}
	inOrder(arr, root.Left, true, lvl+1)
	*arr = append(*arr, root.Val)
	inOrder(arr, root.Right, false, lvl+1)
}

func preOrder(arr *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	preOrder(arr, root.Left)
	preOrder(arr, root.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	inIndex := make([]int, 6001)

	for i, v := range inorder {
		inIndex[v+3000] = i
	}

	var root *TreeNode

	for _, val := range preorder {

		newNode := NewNode(val)

		if root == nil {
			root = newNode
			continue
		}

		valIdx := inIndex[val+3000]
		var parent, current *TreeNode = nil, root
		isLeft := false

		for current != nil {

			currIdx := inIndex[current.Val+3000]
			if valIdx < currIdx {
				parent = current
				current = current.Left
				isLeft = true
			} else {
				parent = current
				current = current.Right
				isLeft = false
			}
		}

		if isLeft {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}

	}

	// arr1 := make([]int, 0)
	// preOrder(&arr1, root)
	// fmt.Println("PREORDER#1", preorder)
	// fmt.Println("PREORDER#2", arr1)

	// arr2 := make([]int, 0)
	// inOrder(&arr2, root, true, 0)
	// fmt.Println("INORDER#1", inorder)
	// fmt.Println("INORDER#2", arr2)

	return root
}
