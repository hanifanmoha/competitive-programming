package findmedianfromdatastream

import "fmt"

type Node struct {
	val    int
	left   *Node
	right  *Node
	nLeft  int
	nRight int
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{val: val}
		return
	}
	t.insertToNode(t.root, val)
}

func (t *Tree) Traverse(node *Node) {
	if node == nil {
		return
	}
	t.Traverse(node.left)
	fmt.Println(node.val)
	t.Traverse(node.right)
}

func (t *Tree) FindMedian() float64 {
	n := 0
	if t.root != nil {
		n = 1 + t.root.nLeft + t.root.nRight
	}

	if n%2 == 1 {
		midIdx := n/2 + 1
		midNode := t.getByIndex(t.root, midIdx)
		return float64(midNode.val)
	} else {
		a, b := n/2, n/2+1
		nodeA, nodeB := t.getByIndex(t.root, a), t.getByIndex(t.root, b)
		return (float64(nodeA.val) + float64(nodeB.val)) / 2
	}
}

func (t *Tree) getByIndex(root *Node, idx int) *Node {
	if root == nil {
		return nil
	}

	if idx <= root.nLeft {
		return t.getByIndex(root.left, idx)
	} else if idx == root.nLeft+1 {
		return root
	} else {
		return t.getByIndex(root.right, idx-root.nLeft-1)
	}
}

func (t *Tree) insertToNode(node *Node, val int) {
	if val < node.val {
		node.nLeft++
		if node.left == nil {
			node.left = &Node{val: val}
		} else {
			t.insertToNode(node.left, val)
		}
	} else {
		node.nRight++
		if node.right == nil {
			node.right = &Node{val: val}
		} else {
			t.insertToNode(node.right, val)
		}
	}
}
