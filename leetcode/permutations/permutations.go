package permutations

func isAllFlagTrue(flag []bool) bool {
	for _, v := range flag {
		if !v {
			return false
		}
	}
	return true
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func permute2(nums []int, flag []bool, temp []int) [][]int {

	if isAllFlagTrue(flag) {
		return [][]int{copySlice(temp)}
	}

	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if flag[i] {
			continue
		}
		flag[i] = true
		newTemp := copySlice(temp)
		newTemp = append(newTemp, nums[i])
		tmpResult := permute2(nums, flag, newTemp)
		result = append(result, tmpResult...)
		flag[i] = false
	}

	return result

}

func permute(nums []int) [][]int {

	flag := make([]bool, len(nums))

	return permute2(nums, flag, []int{})
}
