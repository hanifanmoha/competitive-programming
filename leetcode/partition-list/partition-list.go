package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// In  : [1,4,3,2,5,2]
// Out : [1,2,2,4,3,5]

func print(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append([]int{head.Val}, print(head.Next)...)
}

func partition(head *ListNode, x int) *ListNode {

	var h1, h2, c1, c2 *ListNode

	current := head
	for current != nil {

		if current.Val < x {

			if h1 == nil {
				h1 = current
				c1 = current
			} else {
				c1.Next = current
				c1 = current
			}

		} else {

			if h2 == nil {
				h2 = current
				c2 = current
			} else {
				c2.Next = current
				c2 = current
			}

		}

		tmpNext := current.Next
		current.Next = nil
		current = tmpNext
	}

	// fmt.Println(print(h1))
	// fmt.Println(print(h2))

	if h1 == nil {
		return h2
	}

	c1.Next = h2

	return h1
}
