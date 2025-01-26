package main

func minSubArrayLen(target int, nums []int) int {

	// fmt.Println(nums)

	// 1, 2, 3, 4,  5
	// 0, 3, 5, 7,  9
	// 0, 0, 6, 9,  12
	// 0, 0, 0, 10, 14
	// 0, 0, 0, 0,  15

	n := len(nums)

	tmp := make([]int, n)
	for i := range nums {
		tmp[i] = nums[i]
		if tmp[i] >= target {
			return 1
		}
	}

	for i := 1; i < n; i++ {
		tmp2 := make([]int, n)
		for j := n - 1; j >= i; j-- {
			x := nums[j] + tmp[j-1]
			if x >= target {
				return i + 1
			}
			tmp2[j] = x
		}
		// fmt.Println(i, " - ", tmp2)
		tmp = tmp2
	}

	return 0
}
