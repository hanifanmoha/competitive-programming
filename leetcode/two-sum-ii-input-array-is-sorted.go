package main

// 2 7 11 15

func twoSum(numbers []int, target int) []int {

	n := len(numbers)

	a, b := 0, n-1

	for a < b {
		sum := numbers[a] + numbers[b]
		if sum == target {
			return []int{a + 1, b + 1}
		} else if sum < target {
			a++
		} else {
			b--
		}
	}

	return []int{}
}
