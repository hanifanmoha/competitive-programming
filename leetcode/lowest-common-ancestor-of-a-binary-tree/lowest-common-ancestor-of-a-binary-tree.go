package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	CONTAINS_BOTH = 3
	CONTAINS_P    = 2
	CONTAINS_Q    = 1
	NOT_CONTAINS  = 0
)

func lowestCommenAncestor2(root *TreeNode, p, q int) (int, *TreeNode) {
	if root == nil {
		return NOT_CONTAINS, nil
	}

	leftStatus, lca := lowestCommenAncestor2(root.Left, p, q)
	if leftStatus == CONTAINS_BOTH {
		return leftStatus, lca
	}

	rigthStatus, lca := lowestCommenAncestor2(root.Right, p, q)
	if rigthStatus == CONTAINS_BOTH {
		return rigthStatus, lca
	}

	status := NOT_CONTAINS
	if root.Val == p {
		status = CONTAINS_P
	} else if root.Val == q {
		status = CONTAINS_Q
	}

	status = status | leftStatus | rigthStatus

	if status == CONTAINS_BOTH {
		return status, root
	}

	return status, nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, lca := lowestCommenAncestor2(root, p.Val, q.Val)
	return lca
}
