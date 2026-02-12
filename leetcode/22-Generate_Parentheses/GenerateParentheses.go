package generateparentheses

import "strings"

func generateCombinations(n int) [][]rune {

	stack := make([]rune, 0)
	res := make([][]rune, 0)
	backtrack(n, 0, 0, &stack, &res)
	return res
}

func backtrack(n, openN, closeN int, stack *[]rune, res *[][]rune) {
	if openN == n && closeN == n {
		s := make([]rune, len(*stack))
		copy(s, *stack)
		*res = append(*res, s)
		return
	}

	if openN < n {
		*stack = append(*stack, '(')
		backtrack(n, openN+1, closeN, stack, res)
		*stack = (*stack)[:len(*stack)-1]
	}

	if closeN < openN {
		*stack = append(*stack, ')')
		backtrack(n, openN, closeN+1, stack, res)
		*stack = (*stack)[:len(*stack)-1]
	}
}

// =============================================
// https://neetcode.io/solutions/generate-parentheses
func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	res := make([]string, 0)

	var backtrack func(int, int)
	backtrack = func(openN, closeN int) {
		if openN == n && closeN == n {
			res = append(res, strings.Join(stack, ""))
			return
		}

		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closeN)
			stack = stack[:len(stack)-1]
		}

		if closeN < openN {
			stack = append(stack, ")")
			backtrack(openN, closeN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return res
}
