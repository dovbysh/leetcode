package longestValidParentheses

type be [2]int

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
	index := make([]be, 0, lens/2)
	for i := 0; i < lens-1; i++ {
		if b[i] == '(' && b[i+1] == ')' {
			index = append(index, [2]int{i, i + 1})
			i++
		}
	}
	index = reduceIndex(index)
	for {
		for i := 0; i < len(index); i++ {
			in := index[i]
			if in[0]-1 < 0 {
				continue
			}
			if in[1]+1 >= lens {
				continue
			}
			if b[in[0]-1] == '(' && b[in[1]+1] == ')' {
				index[i][0]--
				index[i][1]++
				i--
				continue
			}
		}
		leni := len(index)
		index = reduceIndex(index)
		if len(index) == leni {
			break
		}
	}
	maxi = 0
	for _, b2 := range index {
		if b2[1]-b2[0]+1 > maxi {
			maxi = b2[1] - b2[0] + 1
		}
	}
	return maxi
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

func reduceIndex(index []be) []be {
	lastI := len(index) - 1
	found := false
	lastFound := 0
	for i := 0; i <= lastI-1; i++ {
		if !found && index[i][1] == (index[i+1][0]-1) {
			found = true
			lastFound = i
			index[i][1] = index[i+1][1]
			i++
			continue
		}
		if found {
			if index[lastFound][1] == (index[i][0] - 1) {
				index[lastFound][1] = index[i][1]
			} else {
				index = append(index[:lastFound+1], index[i:]...)
				found = false
				lastI = len(index) - 1
				i = lastFound
			}
		}
	}
	if found {
		i := lastI
		if index[lastFound][1] == (index[i][1]) {
			index[lastFound][1] = index[i][1]
			index = index[:lastFound+1]
		} else {
			index = append(index[:lastFound+1], index[i:]...)
			found = false
			lastI = len(index) - 1
			i = lastFound
		}
	}
	return index
}
