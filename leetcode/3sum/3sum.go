package three_sum

import "sort"

func sort3(a, b, c int) []int {
	if a <= b && a <= c {
		if b <= c {
			return []int{a, b, c}
		} else {
			return []int{a, c, b}
		}
	} else if b <= a && b <= c {
		if a <= c {
			return []int{b, a, c}
		} else {
			return []int{b, c, a}
		}
	} else {
		if a <= b {
			return []int{c, a, b}
		} else {
			return []int{c, b, a}
		}
	}
}

func insert(myMap map[int]map[int]map[int]bool, vals []int) {
	vals = sort3(vals[0], vals[1], vals[2])

	if _, ok1 := myMap[vals[0]]; !ok1 {
		myMap[vals[0]] = make(map[int]map[int]bool)
	}

	if _, ok2 := myMap[vals[0]][vals[1]]; !ok2 {
		myMap[vals[0]][vals[1]] = make(map[int]bool)
	}

	if _, ok3 := myMap[vals[0]][vals[1]][vals[2]]; !ok3 {
		myMap[vals[0]][vals[1]][vals[2]] = true
	}
}

func threeSum2(nums []int) [][]int {

	n := len(nums)
	sort.Ints(nums)

	tmpRes := make(map[int]map[int]map[int]bool, 0)

	for a := 0; a < n; {
		va := nums[a]
		for b := a + 1; b < n; {
			vb := nums[b]
			for c := b + 1; c < n; {
				vc := nums[c]
				if va+vb+vc == 0 {
					insert(tmpRes, []int{va, vb, vc})
				}
				for c < n && nums[c] == vc {
					c++
				}
			}
			for b < n && nums[b] == vb {
				b++
			}
		}
		for a < n && nums[a] == va {
			a++
		}
	}

	res := make([][]int, 0)

	for v1, m1 := range tmpRes {
		for v2, m2 := range m1 {
			for v3 := range m2 {
				res = append(res, []int{v1, v2, v3})
			}
		}
	}

	return res
}
