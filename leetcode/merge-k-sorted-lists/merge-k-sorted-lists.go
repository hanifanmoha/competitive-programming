package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func list2Arr(p *ListNode) []int {
	arr := make([]int, 0)
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	return arr
}

func isAllPointersNull(pointers []*ListNode) bool {
	for _, p := range pointers {
		if p != nil {
			return false
		}
	}
	return true
}

func mergeKLists(lists []*ListNode) *ListNode {

	n := len(lists)
	pointers := make([]*ListNode, len(lists))
	for i := range lists {
		pointers[i] = lists[i]
		if pointers[i] == nil {
			n--
		}
	}

	var head, tail *ListNode

	// for !isAllPointersNull(pointers) {
	for n > 0 {

		minVal := 999999
		minIdx := 0

		for i := range pointers {
			node := pointers[i]
			if node != nil && node.Val < minVal {
				minVal = node.Val
				minIdx = i
			}

		}

		if head == nil {
			head = pointers[minIdx]
			tail = pointers[minIdx]
		} else {
			tail.Next = pointers[minIdx]
			tail = pointers[minIdx]
		}

		pointers[minIdx] = pointers[minIdx].Next
		if pointers[minIdx] == nil {
			n--
		}

	}

	return head
}
