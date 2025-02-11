package basiccalculator

import (
	"fmt"
	"strconv"
)

type Stack struct {
	stack []string
	idx   int
}

func (this *Stack) Init() {
	this.stack = make([]string, 0)
	this.idx = -1
}

func (this *Stack) Push(s string) {
	this.idx++
	if this.idx < len(this.stack) {
		this.stack[this.idx] = s
	} else {
		this.stack = append(this.stack, s)
	}
}

func (this *Stack) Pop() string {
	tmp := this.stack[this.idx]
	this.idx--
	return tmp
}

func (this *Stack) Top() string {
	return this.stack[this.idx]
}

func (this *Stack) IsEmpty() bool {
	return this.idx == -1
}

func (this *Stack) Vals() []string {
	return this.stack[0 : this.idx+1]
}

func removeSpace(s string) string {
	str := make([]byte, 0)
	for i := range s {
		if s[i] != ' ' {
			str = append(str, s[i])
		}
	}
	return string(str)
}

func operatorOrder(c byte) int {
	if c == '(' || c == ')' {
		return 3
	}
	if c == '^' {
		return 2
	}
	if c == '/' || c == '*' {
		return 1
	}
	if c == '+' || c == '-' {
		return 0
	}
	return 0
}

func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func isParenthesis(c byte) bool {
	return c == '(' || c == ')'
}

// 3 + 8 / 2 * 5
func infix2Postfix(tokens []string) []string {
	stack := (Stack{})
	stack.Init()

	postfix := make([]string, 0)
	for _, token := range tokens {
		if isOperator(token[0]) {
			if stack.IsEmpty() {
				stack.Push(token)
			} else {
				top := stack.Top()
				// fmt.Println("Operator collision", string(token[0]), string(top[0]), operatorOrder(token[0]), operatorOrder(top[0]))
				if isParenthesis(top[0]) || operatorOrder(token[0]) > operatorOrder(top[0]) {
					stack.Push(token)
				} else {
					// fmt.Println("Move", top, "to result")
					stack.Pop()
					stack.Push(token)
					postfix = append(postfix, top)
				}
			}
		} else if isParenthesis(token[0]) {
			if token[0] == '(' {
				// fmt.Println("Stack (", o)
				stack.Push(token)
			} else {
				tmp := ""
				for tmp != "(" {
					tmp = stack.Pop()
					// fmt.Println("Unstack", tmp, o)
					// fmt.Println("Current Stack", stack.Vals())
					if tmp != "(" {
						postfix = append(postfix, tmp)
					}
				}
			}
		} else {
			postfix = append(postfix, token)
		}

	}
	for !stack.IsEmpty() {
		postfix = append(postfix, stack.Pop())
	}
	return postfix
}

func str2Token(s string) []string {
	tokens := make([]string, 0)
	tmp := make([]byte, 0)
	isAfterNumber := false
	for i := range s {
		if s[i] == '-' && !isAfterNumber {
			// tokens = append(tokens, "~")
			tmp = make([]byte, 0)
		} else if isOperator(s[i]) || isParenthesis(s[i]) {
			if len(tmp) > 0 {
				tokens = append(tokens, string(tmp))
			}
			tokens = append(tokens, s[i:i+1])
			tmp = make([]byte, 0)
			if isOperator(s[i]) {
				isAfterNumber = false
			}
		} else {
			tmp = append(tmp, s[i])
			isAfterNumber = true
		}
	}
	if len(tmp) > 0 {
		tokens = append(tokens, string(tmp))
	}
	return tokens
}

func evaluatePostfix(tokens []string) int {

	stack := Stack{}
	stack.Init()

	for _, t := range tokens {
		if len(t) == 1 && isOperator(t[0]) {
			b, _ := strconv.Atoi(stack.Pop())
			a, _ := strconv.Atoi(stack.Pop())
			tmp := 0
			if t[0] == '+' {
				tmp = a + b
			} else if t[0] == '-' {
				tmp = a - b
			} else if t[0] == '*' {
				tmp = a * b
			} else if t[0] == '/' {
				tmp = a / b
			}
			// fmt.Println("Evaluate", a, t, b, "=", tmp)
			stack.Push(fmt.Sprint(tmp))
		} else {
			stack.Push(t)
		}
	}

	finalRes, _ := strconv.Atoi(stack.Top())

	return finalRes
}

func calculate(s string) int {
	// fmt.Println("Original Problem", s)
	s = removeSpace(s)
	// fmt.Println("Without Spaces", s)
	tokens := str2Token(s)
	fmt.Println("Tokens", tokens)
	postfix := infix2Postfix(tokens)
	fmt.Println(postfix)
	return evaluatePostfix(postfix)
}

/*

1 + ( 2 + 3 )

*/
