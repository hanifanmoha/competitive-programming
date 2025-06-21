package editdistance

func minof2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minof3(a, b, c int) int {
	return minof2(a, minof2(b, c))
}

func minDistanceIdx(str, target string, idx1, idx2 int, mem [][]*int) int {
	n1, n2 := len(str), len(target)
	if idx1 >= n1 && idx2 >= n2 {
		return 0
	} else if idx1 >= n1 {
		// str : abc; target : abc123; -> insertion
		return n2 - idx2
	} else if idx2 >= n2 {
		// str : abc123; target : abc; -> deletion
		return n1 - idx1
	}

	if mem[idx1][idx2] != nil {
		return *mem[idx1][idx2]
	}

	res := 0
	if str[idx1] == target[idx2] {
		res = minof3(
			minDistanceIdx(str, target, idx1+1, idx2+1, mem),
			1+minDistanceIdx(str, target, idx1, idx2+1, mem), // insertion
			1+minDistanceIdx(str, target, idx1+1, idx2, mem), // deletion
		)
	} else {
		res = minof3(
			1+minDistanceIdx(str, target, idx1+1, idx2+1, mem), //replace
			1+minDistanceIdx(str, target, idx1, idx2+1, mem),   // insertion
			1+minDistanceIdx(str, target, idx1+1, idx2, mem),   // deletion
		)
	}
	mem[idx1][idx2] = &res
	return res
}

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	mem := make([][]*int, n1)
	for i := range n1 {
		mem[i] = make([]*int, n2)
	}

	return minDistanceIdx(word1, word2, 0, 0, mem)
}
