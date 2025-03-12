package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func countNode(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + countNode(head.Next)
}

func cutNode(head *ListNode, count int) *ListNode {
	if head == nil {
		return nil
	}

	if count == 1 {
		next := head.Next
		head.Next = nil
		return next
	}

	return cutNode(head.Next, count-1)
}

func list2arr(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	tmp := list2arr(head.Next)
	return append([]int{head.Val}, tmp...)
}

func sortList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	count := countNode(head)
	if count == 1 {
		return head
	}

	mid := count / 2
	midNode := cutNode(head, mid)

	a := sortList(head)
	b := sortList(midNode)

	c1, c2 := a, b

	var h *ListNode
	var t *ListNode

	for c1 != nil && c2 != nil {
		if c1.Val < c2.Val {
			if h == nil {
				h = c1
				t = c1
			} else {
				t.Next = c1
			}
			t = c1
			c1 = c1.Next
		} else {
			if h == nil {
				h = c2
				t = c2
			} else {
				t.Next = c2
			}
			t = c2
			c2 = c2.Next
		}
	}

	for c1 != nil {
		if h == nil {
			h = c1
			t = c1
		} else {
			t.Next = c1
		}
		t = c1
		c1 = c1.Next
	}

	for c2 != nil {
		if h == nil {
			h = c2
			t = c2
		} else {
			t.Next = c2
		}
		t = c2
		c2 = c2.Next
	}

	return h
}
