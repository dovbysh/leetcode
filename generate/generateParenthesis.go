package generate

import (
	"fmt"
	"sort"
	"strings"
)

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	res := make([]string, 0, n)
	bitts := n << 1
	mask := "%0" + fmt.Sprint(bitts) + "b"
	maxI := 1 << bitts
	ex := make(map[string]struct{}, n)
	exists := false
	for i := 0; i < maxI; i++ {
		z := fmt.Sprintf(mask, i)
		zz := strings.Replace(
			strings.Replace(z, "0", "(", -1),
			"1", ")", -1)
		if check(zz) {
			if _, exists = ex[zz]; exists {
				continue
			}
			ex[zz] = struct{}{}
			res = append(res, zz)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
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
