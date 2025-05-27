package binarytreemaximumpathsum

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mps2(root *TreeNode) (maxSumDiscontinue int, maxSumContinue int) {

	if root == nil {
		return -999999999, 0
	}

	msdLeft, mscLeft := mps2(root.Left)
	msdRight, mscRight := mps2(root.Right)

	maxSumContinue = getMax(root.Val, getMax(root.Val+mscLeft, root.Val+mscRight))

	maxSumDiscontinue = getMax(root.Val, getMax(msdLeft, getMax(msdRight, getMax(root.Val+mscLeft, getMax(root.Val+mscRight, root.Val+mscLeft+mscRight)))))

	return
}

func maxPathSum(root *TreeNode) int {

	res, _ := mps2(root)

	return res
}
