package substring_with_concatenation_of_all_words

import "fmt"

type WordCounter struct {
	// validWordCount int
	wordCount map[string]int

	// tmpValidWordCount int
	tmpCount map[string]int
}

func (this *WordCounter) Init(words []string) {
	this.wordCount = make(map[string]int)
	for _, word := range words {
		// this.validWordCount++
		if _, exist := this.wordCount[word]; exist {
			this.wordCount[word]++
		} else {
			this.wordCount[word] = 1
		}
	}
}

func (this *WordCounter) Reset() {
	this.tmpCount = make(map[string]int)
	// this.tmpValidWordCount = this.validWordCount
	for word, count := range this.wordCount {
		this.tmpCount[word] = count
	}
}

func (this *WordCounter) Add(word string) {
	if _, isValidWord := this.wordCount[word]; !isValidWord {
		// this.tmpValidWordCount++
		return
	}

	if _, isExist := this.tmpCount[word]; isExist {
		this.tmpCount[word] += 1
	} else {
		this.tmpCount[word] = 1
	}
}

func (this *WordCounter) Remove(word string) {
	if _, isValidWord := this.wordCount[word]; !isValidWord {
		// this.tmpValidWordCount--
		return
	}

	if _, isExist := this.tmpCount[word]; isExist {
		this.tmpCount[word] -= 1
	}
}

func (this *WordCounter) IsValid() bool {
	// return this.tmpValidWordCount == 0
	for _, count := range this.tmpCount {
		if count != 0 {
			return false
		}
	}
	// fmt.Println("Valid")
	return true
}

func findSubstring(s string, words []string) []int {

	fmt.Printf("S : %s\n", s)
	fmt.Printf("Words : %v\n", words)

	nW := len(words)
	lenS := len(s)

	if nW <= 0 || lenS <= 0 {
		return []int{}
	}

	lenW := len(words[0])

	if lenW <= 0 {
		return []int{}
	}

	indexes := make([]int, 0)

	wc := WordCounter{}
	wc.Init(words)

	for i := 0; i < lenW; i++ {
		wc.Reset()
		for j := i; j+lenW <= lenS; j += lenW {
			newWord := s[j : j+lenW]
			wc.Remove(newWord)
			// fmt.Printf("%s ", newWord)
			startIndex := j - (nW-1)*lenW
			if startIndex-lenW >= 0 {
				oldWord := s[startIndex-lenW : startIndex]
				// fmt.Println(startIndex, "Remove", newWord, "Add", oldWord)
				wc.Add(oldWord)
			} else {
				// fmt.Println(startIndex, "Remove", newWord)
			}
			if wc.IsValid() {
				indexes = append(indexes, startIndex)
			}
		}
		// fmt.Println()
	}

	return indexes
}
