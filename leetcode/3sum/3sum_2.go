package three_sum

const OFFSET = 200000

type Mem struct {
	mem []bool
}

func NewMem() *Mem {
	return &Mem{
		mem: make([]bool, 1000000),
	}
}

func (m *Mem) Insert(val int) {
	idx := val + OFFSET
	m.mem[idx] = true
}

func (m *Mem) IsExist(val int) bool {
	idx := val + OFFSET
	return m.mem[idx]
}

func threeSum(nums []int) [][]int {

	mem := NewMem()

	n := len(nums)

	tmpRes := make(map[int]map[int]map[int]bool, 0)

	// for _, val := range nums {
	// 	mem.Insert(val)
	// }

	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			sum2 := nums[a] + nums[b]
			if mem.IsExist(-sum2) {
				insert(tmpRes, []int{-sum2, nums[a], nums[b]})
			}
		}

		mem.Insert(nums[a])
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
