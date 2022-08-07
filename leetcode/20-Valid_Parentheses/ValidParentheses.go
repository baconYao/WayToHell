package validparentheses

func ValidParentheses(s string) bool {
	stack := []string{}
	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, string(')'))
		case '{':
			stack = append(stack, string('}'))
		case '[':
			stack = append(stack, string(']'))
		default:
			if len(stack) == 0 || stack[len(stack)-1] != string(char) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
