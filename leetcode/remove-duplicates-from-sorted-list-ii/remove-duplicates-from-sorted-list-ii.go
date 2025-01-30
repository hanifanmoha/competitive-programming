package remove_duplicates_from_sorted_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

// In  : [1,2,3,3,4,4,5]
// Out : [1,2,5]

func deleteDuplicates(head *ListNode) *ListNode {

	var newHead, tail, prev *ListNode

	current := head
	for current != nil {

		isSimilarWithPrev := prev != nil && prev.Val == current.Val
		isSimilarWithNext := current.Next != nil && current.Next.Val == current.Val

		if !isSimilarWithNext && !isSimilarWithPrev {

			if newHead == nil {
				newHead = current
				tail = current
			} else {
				tail.Next = current
				tail = current
			}

		}

		prev = current
		tmpNext := current.Next
		current.Next = nil
		current = tmpNext
	}

	return newHead
}
