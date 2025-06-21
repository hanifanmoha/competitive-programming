package maximumsumcircularsubarray

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSubarraySumCircular(nums []int) int {

	sum, minSum, maxSum := 0, 9999999, -9999999
	tmpMinSum, tmpMaxSum := 9999999, -9999999

	for _, n := range nums {
		tmpMinSum = min(tmpMinSum+n, n)
		tmpMaxSum = max(tmpMaxSum+n, n)

		// fmt.Println(i, n, tmpMinSum, tmpMaxSum)

		sum += n
		minSum = min(minSum, tmpMinSum)
		maxSum = max(maxSum, tmpMaxSum)
	}

	// fmt.Println(sum, maxSum, minSum)

	if sum == minSum {
		return maxSum
	}

	return max(maxSum, sum-minSum)
}
