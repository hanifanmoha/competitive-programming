package houserobber

func rob2(nums []int, mem []int, start int) int {
	n := len(nums)
	ln := n - start

	if start >= n {
		return 0
	}

	if ln == 1 {
		return nums[start]
	}

	if mem[start] >= 0 {
		return mem[start]
	}

	o1 := nums[start] + rob2(nums, mem, start+2)
	o2 := rob2(nums, mem, start+1)

	if o1 > o2 {
		mem[start] = o1
	} else {
		mem[start] = o2
	}

	return mem[start]
}

func rob(nums []int) int {
	n := len(nums)

	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}

	return rob2(nums, mem, 0)
}
