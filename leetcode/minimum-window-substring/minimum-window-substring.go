package minimum_window_substring

import "fmt"

func c2i(c rune) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A')
	}
	return int(c-'a') + 26
}

type Counter struct {
	counter []int
}

func (this *Counter) Init(t string) {
	this.counter = make([]int, 52)
	for _, c := range t {
		this.counter[c2i(c)]++
	}
}

func (this *Counter) Add(c rune) {
	fmt.Println("ADD", string(c), "to", this.counter[c2i(c)]+1)
	this.counter[c2i(c)]++
}

func (this *Counter) Reduce(c rune) {
	fmt.Println("REDUCE", string(c), "to", this.counter[c2i(c)]-1)
	this.counter[c2i(c)]--
}

func (this *Counter) IsValid() bool {
	for idx, count := range this.counter {
		if count > 0 {
			fmt.Println(idx, "causes invalid with count", count)
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	fmt.Println("S:", s, "T:", t)

	counter := Counter{}
	counter.Init(t)

	ls := len(s)
	minLength := 999999
	str := ""

	left, right := 0, -1

	for left < ls && right < ls {

		fmt.Println("CHECK", left, right)

		if counter.IsValid() {

			fmt.Println("VALID", s[left:right+1])

			newLen := right - left + 1
			if newLen < minLength {
				minLength = newLen
				str = s[left : right+1]
			}

			counter.Add(rune(s[left]))
			left++
		} else {

			fmt.Println("INVALID", s[left:right+1])
			right++
			if right < ls {
				counter.Reduce(rune(s[right]))
			}
		}
	}

	return str
}
