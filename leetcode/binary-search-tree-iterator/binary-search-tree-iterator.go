package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	Data  []*TreeNode
	index int
}

func StackConstructor() Stack {
	stack := Stack{}
	stack.Data = make([]*TreeNode, 0)
	stack.index = -1
	return stack
}

func (this *Stack) Push(node *TreeNode) {
	l := len(this.Data)
	this.index++
	if this.index < l {
		this.Data[this.index] = node
	} else {
		this.Data = append(this.Data, node)
	}
	return
}

func (this *Stack) Pop() {
	this.index--
	return
}

func (this *Stack) Top() *TreeNode {
	return this.Data[this.index]
}

type BSTIterator struct {
	root    *TreeNode
	current *TreeNode
	parents *Stack
}

func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{}
	bstIterator.current = nil
	bstIterator.root = root
	stack := StackConstructor()
	bstIterator.parents = &stack
	return bstIterator
}

func (this *BSTIterator) Next() int {

	if this.current == nil {
		this.current = this.root
		for this.current.Left != nil {
			this.parents.Push(this.current)
			this.current = this.current.Left
		}
		return this.current.Val
	}

	isLeftChild := this.current == this.parents.Top().Left

	if isLeftChild {
		// up to parent that has right
		// c := this.parents.Top()

		// go to right
		// go left, push to stack
	}

	return this.current.Val
}

func (this *BSTIterator) HasNext() bool {
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
