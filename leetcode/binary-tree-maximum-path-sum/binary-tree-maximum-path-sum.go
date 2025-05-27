package binarytreemaximumpathsum

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type GraphNode struct {
	Val       int
	Adjacent  []*GraphNode
	FlagStart bool
	FlagPath  bool
}

const MAX_VAL int = 999999999

func NewGraphNode(Val int) *GraphNode {
	n := GraphNode{
		Val:       Val,
		Adjacent:  make([]*GraphNode, 0),
		FlagStart: false,
		FlagPath:  false,
	}
	return &n
}

func traverse(root *TreeNode) *GraphNode {
	if root == nil {
		return nil
	}

	gNode := NewGraphNode(root.Val)

	gLeft := traverse(root.Left)
	if gLeft != nil {
		gNode.Adjacent = append(gNode.Adjacent, gLeft)
		gLeft.Adjacent = append(gLeft.Adjacent, gNode)
	}

	gRight := traverse(root.Right)
	if gRight != nil {
		gNode.Adjacent = append(gNode.Adjacent, gRight)
		gRight.Adjacent = append(gRight.Adjacent, gNode)
	}

	return gNode
}

func findStartPath(startNode *GraphNode) {
	if startNode.FlagStart {
		return
	}

	fmt.Println("Attemp to start from ", startNode.Val)

	startNode.FlagStart = true
	for _, s := range startNode.Adjacent {
		findStartPath(s)
	}

	tmpMaxPath := calcPath(startNode)
	if tmpMaxPath > mps {
		mps = tmpMaxPath
	}
}

func calcPath(current *GraphNode) int {
	// flag as true
	// iterate from adjacent, get calc from there
	// get maximum
	// flag as false
	// return val + tot

	if current.FlagPath {
		return -MAX_VAL
	}

	current.FlagPath = true

	maxTot := current.Val
	for _, next := range current.Adjacent {
		tot := calcPath(next) + current.Val
		if tot > maxTot {
			maxTot = tot
		}
	}

	current.FlagPath = false

	return maxTot
}

var mps int = -MAX_VAL

func maxPathSum2(root *TreeNode) int {

	mps = -MAX_VAL

	gRoot := traverse(root)

	findStartPath(gRoot)

	return mps
}
