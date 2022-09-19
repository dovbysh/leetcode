package longestValidParentheses

func longestValidParentheses(s string) int {
	return 0
}

func check(zz string) bool {
	c := 0
	for _, r := range zz {
		switch r {
		case '(':
			c++
			break
		case ')':
			c--
			break
		}
		if c < 0 {
			return false
		}
	}
	return c == 0
}
