package three_sum

import "sort"

func threeSum(nums []int) [][]int {

	n := len(nums)
	sort.Ints(nums)

	res := make([][]int, 0)

	// va, vb, vc := 999999, 999999

	for a := 0; a < n; {
		va := nums[a]
		for b := a + 1; b < n; {
			vb := nums[b]
			for c := b + 1; c < n; {
				vc := nums[c]

				if va+vb+vc == 0 {
					res = append(res, []int{va, vb, vc})
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

	return res
}
