package bitwiseandofnumbersrange

func rangeBitwiseAnd(left int, right int) int {

	n := 0
	for left != right {
		left /= 2
		right /= 2
		n++
	}

	for range n {
		left *= 2
	}

	return left
}
