package lettercombinationsofaphonenumber

func initDigits() map[rune][]rune {
	return map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
}

func letterCombinations(digits string) []string {

	digitsMap := initDigits()

	n := len(digits)
	res := make([][]rune, 0)
	for i := n - 1; i >= 0; i-- {

		chars := digitsMap[rune(digits[i])]
		newRes := make([][]rune, 0)

		if len(res) == 0 {
			for _, c := range chars {
				newRes = append(newRes, []rune{c})
			}
		} else {
			for _, c := range chars {
				for _, r := range res {
					newRes = append(newRes, append([]rune{c}, r...))
				}
			}
		}
		res = newRes
	}

	resString := make([]string, 0)
	for _, r := range res {
		resString = append(resString, string(r))
	}

	return resString
}
