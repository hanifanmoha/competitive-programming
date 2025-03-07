package coinchange

func mergeSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	a := arr[:n/2]
	b := arr[n/2:]

	a = mergeSort(a)
	b = mergeSort(b)

	nA, nB := len(a), len(b)
	idxA, idxB := 0, 0
	res := make([]int, 0)
	for idxA < nA && idxB < nB {
		if a[idxA] > b[idxB] {
			res = append(res, a[idxA])
			idxA++
		} else {
			res = append(res, b[idxB])
			idxB++
		}
	}
	for idxA < nA {
		res = append(res, a[idxA])
		idxA++
	}
	for idxB < nB {
		res = append(res, b[idxB])
		idxB++
	}
	return res
}

const UNINITIALIZED = -2
const INVALID_VAL = -1
const MAX_VAL = 999999

func coinChange(coins []int, amount int) int {
	mem := make([]int, amount+1)
	for i := range mem {
		mem[i] = UNINITIALIZED
	}
	return coinChange2(coins, amount, mem)
}

func coinChange2(coins []int, amount int, mem []int) int {

	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	if mem[amount] != UNINITIALIZED {
		return mem[amount]
	}

	coins = mergeSort(coins)

	minVal := MAX_VAL
	for _, c := range coins {
		cc := coinChange2(coins, amount-c, mem)
		if cc != INVALID_VAL && cc < minVal {
			minVal = cc
		}
	}

	if minVal < MAX_VAL {
		mem[amount] = minVal + 1
		return minVal + 1
	}

	mem[amount] = INVALID_VAL
	return INVALID_VAL
}
