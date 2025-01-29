package text_justification

import "strings"

func getMaximize(maxWidth int, start int, words []string) int {
	n := len(words)
	end := start
	cnt := 0
	for idx := start; idx < n; idx++ {
		numSpaces := idx - start
		if cnt+len(words[idx])+numSpaces <= maxWidth {
			cnt += len(words[idx])
			end = idx
		} else {
			break
		}
	}
	return end
}

func justify(width int, words []string, isLastRow bool) string {

	n := len(words)
	res := ""

	originalLen := 0
	for _, word := range words {
		originalLen += len(word)
	}

	spaceNeeded := width - originalLen

	// Single word or last row
	if n == 1 || isLastRow {
		res = strings.Join(words, " ")
		for len(res) < width {
			res += " "
		}
		return res
	}

	// Normal row

	defaultSpace := spaceNeeded / (n - 1)
	additionalSpace := spaceNeeded % (n - 1)

	for idx, word := range words {
		res += word

		// Add space
		if idx < n-1 {

			// Default space
			for i := 0; i < defaultSpace; i++ {
				res += " "
			}

			if idx < additionalSpace {
				res += " "
			}
		}
	}

	return res
}

func fullJustify(words []string, maxWidth int) []string {

	// get maximize

	N := len(words)
	start := 0
	list := make([][]string, 0)

	for start < N {
		end := getMaximize(maxWidth, start, words)
		if start <= end {
			tmp := words[start : end+1]
			list = append(list, tmp)
			start = end + 1
		}
	}

	// align
	list2 := make([]string, 0)
	nList := len(list)
	for idx, row := range list {
		list2 = append(list2, justify(maxWidth, row, idx == nList-1))
	}

	return list2
}
