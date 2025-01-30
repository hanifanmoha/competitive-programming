package longest_consecutive_sequence

func calc(mymap map[int]int, key int) int {

	// if not exist, 0
	ln, exist := mymap[key]
	if !exist {
		return 0
	}

	// if > 0, return val
	if ln > 0 {
		return ln
	}

	// return next + 1
	newLn := 1 + calc(mymap, key+1)
	mymap[key] = newLn
	return newLn
}

func longestConsecutive2(nums []int) int {

	mymap := make(map[int]int)

	for _, val := range nums {
		mymap[val] = 0
	}

	// get max
	max := 0
	for key := range mymap {
		ln := calc(mymap, key)
		if ln > max {
			max = ln
		}
	}

	return max
}

func longestConsecutive(nums []int) int {

	mymap := make(map[int]int)

	for _, val := range nums {
		mymap[val] = 0
	}

	// get max
	max := 0
	for key := range mymap {
		if _, prevExist := mymap[key-1]; prevExist {
			continue
		}
		tmp := key
		flag := true
		for flag {
			_, flag = mymap[tmp+1]
			if flag {
				tmp++
			}
		}
		ln := tmp - key + 1
		if ln > max {
			max = ln
		}
	}

	return max
}
