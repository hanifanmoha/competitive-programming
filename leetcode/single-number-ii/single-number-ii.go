package singlenumberii

func singleNumber(nums []int) int {
	myMap := make(map[int]byte)
	for _, n := range nums {
		if _, exist := myMap[n]; exist {
			myMap[n]++
		} else {
			myMap[n] = 1
		}
	}
	for k, v := range myMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
