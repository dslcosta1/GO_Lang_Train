func isValid(s string) bool {
	var mapParentheses = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, parenthese := range s {
		openParenthese, ok := mapParentheses[parenthese]

		if !ok {
			stack = append(stack, parenthese)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != openParenthese {
				return false
			}
		}
	}

	return len(stack) == 0
}