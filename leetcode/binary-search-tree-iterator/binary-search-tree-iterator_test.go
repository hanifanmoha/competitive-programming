package binary_search_tree_iterator

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := StackConstructor()
	node := &TreeNode{Val: 1}
	stack.Push(node)

	if stack.index != 0 {
		t.Errorf("Expected stack index to be 0, got %d", stack.index)
	}

	if len(stack.Data) != 1 {
		t.Errorf("Expected stack length to be 1, got %d", len(stack.Data))
	}

	if stack.Data[0] != node {
		t.Errorf("Expected stack data to be %v, got %v", node, stack.Data[0])
	}
}

func TestStackPop(t *testing.T) {
	stack := StackConstructor()
	node := &TreeNode{Val: 1}
	stack.Push(node)
	stack.Pop()

	if stack.index != -1 {
		t.Errorf("Expected stack index to be -1, got %d", stack.index)
	}

	if len(stack.Data) != 1 {
		t.Errorf("Expected stack length to be 1, got %d", len(stack.Data))
	}

	if stack.Data[0] != node {
		t.Errorf("Expected stack data to be %v, got %v", node, stack.Data[0])
	}
}
