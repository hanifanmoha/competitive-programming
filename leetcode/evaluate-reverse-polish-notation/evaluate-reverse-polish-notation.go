package evaluate_reverse_polish_notation

import (
	"fmt"
	"strconv"
)

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

type Token struct {
	isValue bool
	value   string
	a       *Token
	b       *Token
}

func (t *Token) Parse(idx int, tokens []string) int {
	t.value = tokens[idx]
	if !isOperator(t.value) {
		t.isValue = true
		return 1
	}

	t.b = &Token{}
	lenB := t.b.Parse(idx+1, tokens)

	t.a = &Token{}
	lenA := t.a.Parse(idx+1+lenB, tokens)

	return 1 + lenA + lenB
}

func (t *Token) Print() string {
	if t.isValue {
		return t.value
	}
	return fmt.Sprintf("(%s %s %s)", t.a.Print(), t.value, t.b.Print())
}

func (t *Token) Evaluate() int {
	if t.isValue {
		num, _ := strconv.Atoi(t.value)
		return num
	}

	a := t.a.Evaluate()
	b := t.b.Evaluate()

	switch t.value {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func reverse(tokens []string) {
	a, b := 0, len(tokens)-1
	for a < b {
		tmp := tokens[a]
		tokens[a] = tokens[b]
		tokens[b] = tmp
		a++
		b--
	}
}

func evalRPN(tokens []string) int {
	// fmt.Println("Tokens:", tokens)
	reverse(tokens)
	// fmt.Println("Reversed:", tokens)
	root := &Token{}
	root.Parse(0, tokens)
	// str := root.Print()
	// fmt.Println("String", str)
	return root.Evaluate()
}
