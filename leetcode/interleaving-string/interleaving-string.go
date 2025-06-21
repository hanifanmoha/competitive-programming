package interleavingstring

func iterate(s1, s2, s3 string, d1, d2, d3, n1, n2, n3 int, mem [][][]*bool) bool {

	if d1 >= n1 && d2 >= n2 && d3 >= n3 {
		return true
	}

	if mem[d1][d2][d3] != nil {
		return *mem[d1][d2][d3]
	}

	trueVal, falseVal := true, false

	if d1 < n1 && d3 < n3 && s1[d1] == s3[d3] && iterate(s1, s2, s3, d1+1, d2, d3+1, n1, n2, n3, mem) {
		mem[d1][d2][d3] = &trueVal
		return true
	}
	if d2 < n2 && d3 < n3 && s2[d2] == s3[d3] && iterate(s1, s2, s3, d1, d2+1, d3+1, n1, n2, n3, mem) {
		mem[d1][d2][d3] = &trueVal
		return true
	}
	mem[d1][d2][d3] = &falseVal
	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)

	mem := make([][][]*bool, n1+1)
	for i := range n1 + 1 {
		mem[i] = make([][]*bool, n2+1)
		for j := range n2 + 1 {
			mem[i][j] = make([]*bool, n3+1)
		}
	}

	return iterate(s1, s2, s3, 0, 0, 0, n1, n2, n3, mem)
}
