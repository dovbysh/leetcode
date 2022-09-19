package longestValidParentheses

func longestValidParentheses(s string) int {
	if check([]byte(s), 0, len(s)-1) {
		return len(s)
	}
	lens := len(s)
	maxi := lens - 2
	if maxi < 1 {
		return 0
	}
	b := []byte(s)
	for i := 1; i <= maxi; i++ {
		if (lens-i)%2 == 1 {
			continue
		}
		for j := 0; j <= i; j++ {
			if check(b, i-j, lens-i-1) {
				return lens - i
			}
		}
	}
	return 0
}

func check(zz []byte, first, length int) bool {
	c := 0
	for i := first; i <= first+length; i++ {
		switch zz[i] {
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
		if c > (first + length - i + 3) {
			return false
		}
	}
	return c == 0
}
