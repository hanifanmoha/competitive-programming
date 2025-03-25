package combinationsum

import "sort"

func getSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func copySlice(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

func cs2(candidates []int, start int, target int, arr []int) [][]int {

	sum := getSum(arr)
	if sum > target {
		return [][]int{}
	}

	if sum == target {
		newArr := copySlice(arr)
		return [][]int{newArr}
	}

	res := make([][]int, 0)

	for i := start; i < len(candidates); i++ {
		newArr := copySlice(arr)
		newArr = append(newArr, candidates[i])
		tmpRes := cs2(candidates, i, target, newArr)
		res = append(res, tmpRes...)
	}

	return res
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return cs2(candidates, 0, target, []int{})
}
