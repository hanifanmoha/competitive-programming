package searchinrotatedsortedarray

// function to search lowest value in a shifted sorted array
func searchStart(nums []int) int {
	// fmt.Println(nums)
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		// fmt.Println(left, mid, right)
		if left == right || nums[left] < nums[right] {
			return left
		} else if nums[left] > nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func getRealIndex(startIdx, idx, n int) int {
	return (startIdx + idx) % n
}

func search(nums []int, target int) int {

	// fmt.Println(nums)
	// fmt.Println("target: ", target)

	startIdx := searchStart(nums)

	// fmt.Println("start index:", startIdx)

	n := len(nums)

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		midRealIdx := getRealIndex(startIdx, mid, n)
		val := nums[midRealIdx]
		if val == target {
			return midRealIdx
		} else if target < val {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1
}
