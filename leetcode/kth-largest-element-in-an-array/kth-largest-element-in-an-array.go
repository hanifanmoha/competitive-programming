package kthlargestelementinanarray

func findKthLargest(nums []int, k int) int {

	heap := NewHeap()

	for i := range nums {
		heap.Insert(nums[i])
	}

	tmp := 0
	for i := 0; i < k; i++ {
		tmp = heap.Pop()
	}

	return tmp
}
