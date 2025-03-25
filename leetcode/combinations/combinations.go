package combinations

var tmp map[int]map[int][][]int

func set(val [][]int, n, k int) {
	if tmp == nil {
		tmp = make(map[int]map[int][][]int)
	}
	if tmp[n] == nil {
		tmp[n] = make(map[int][][]int)
	}
	if tmp[n][k] == nil {
		tmp[n][k] = make([][]int, 0)
	}

	tmp[n][k] = val
}

func get(n, k int) [][]int {
	if tmp == nil {
		return nil
	}
	if tmp[n] == nil {
		return nil
	}
	return tmp[n][k]
}

func combine(n int, k int) [][]int {

	cache := get(n, k)
	if cache != nil {
		return cache
	}

	if n == 0 || k == 0 {
		res := make([][]int, 0)
		set(res, n, k)
		return res
	}

	if k == 1 {
		res := make([][]int, 0)
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		set(res, n, k)
		return res
	}

	// combine(n-1, k)

	res := combine(n-1, k)

	// combine(n-1, k-1) + k
	comb2 := combine(n-1, k-1)
	for _, c := range comb2 {
		c2 := make([]int, len(c)+1)
		copy(c2, c)
		c2[len(c)] = n
		res = append(res, c2)
	}

	set(res, n, k)
	return res
}
