package evaluatereversepolishnotation

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := new(Stack)

	for _, v := range tokens {
		if !isOperator(v) {
			num, _ := strconv.Atoi(v)
			stack.Push(num)
		} else {
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			c := 0
			if v == "+" {
				c = b + a
			} else if v == "-" {
				c = b - a
			} else if v == "*" {
				c = b * a
			} else {
				c = b / a
			}
			stack.Push(c)
		}
	}
	target, _ := stack.Pop()
	return target
}

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem, true
}

func isOperator(operator string) bool {
	operators := []string{"+", "-", "*", "/"}
	for _, v := range operators {
		if operator == v {
			return true
		}
	}
	return false
}
