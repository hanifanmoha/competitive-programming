package wordbreak

const (
	IS_VALID     = 1
	IS_NOT_VALID = 2
)

func isSame(s1, s2 string) bool {
	n := len(s1)
	if len(s2) != n {
		return false
	}
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func isInWordDict(s string, start int, to int, wordDict []string) bool {
	for _, w := range wordDict {
		if isSame(w, s[start:to]) {
			return true
		}
	}
	return false
}

func rec(s string, mem []int, start int, wordDict []string) bool {
	n := len(s)
	if start >= n {
		return true
	}

	if mem[start] > 0 {
		return mem[start] == IS_VALID
	}

	for to := start + 1; to <= n; to++ {
		a := isInWordDict(s, start, to, wordDict)
		b := rec(s, mem, to, wordDict)
		if a && b {
			mem[start] = IS_VALID
			return true
		}
	}

	mem[start] = IS_NOT_VALID
	return false
}

func wordBreak(s string, wordDict []string) bool {

	n := len(s)
	mem := make([]int, n)

	return rec(s, mem, 0, wordDict)
}
