package reverse_word_in_a_string

import (
	"strings"
)

func reverseWords(s string) string {

	list := make([]string, 0)

	start := 0
	for idx, c := range s {
		if c == ' ' {
			if idx > start {
				tmp := s[start:idx]
				list = append(list, tmp)
				// fmt.Println(tmp)
			}
			start = idx + 1
		}
	}

	if len(s) > start {
		tmp := s[start:]
		list = append(list, tmp)
		// fmt.Println(tmp)
	}

	// fmt.Println(list)

	n := len(list)
	for i := 0; i < n/2; i++ {
		list[i], list[n-1-i] = list[n-1-i], list[i]
	}

	return strings.Join(list, " ")
}
