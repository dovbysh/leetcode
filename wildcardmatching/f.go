package wildcardmatching

func isMatch(s string, p string) bool {
	if p == "*" {
		return true
	}

	z := z{
		memo:  make([]byte, (len(s)+2)*(len(p)+2)),
		len_s: len(s),
		len_p: len(p),
	}
	return z.dp(0, 0, s, p)
}

type z struct {
	memo         []byte
	len_s, len_p int
}

func (z *z) dp(i, j int, s, p string) bool {
	k := i*(z.len_p+1) + j
	if v := z.memo[k]; v > 0 {
		return v == 2
	}
	if j == z.len_p {
		v := i == z.len_s
		z.memo[k] = 1
		if v {
			z.memo[k] = 2
		}
		return v
	}
	if i == z.len_s {
		for ; j < z.len_p; j++ {
			if p[j] != '*' {
				z.memo[k] = 1
				return false
			}
		}
		z.memo[k] = 2
		return true
	}
	ans := false
	if p[j] == '*' {
		ans = z.dp(i, j+1, s, p) || z.dp(i+1, j, s, p)
	} else {
		match := s != "" && (p[j] == s[i] || p[j] == '?')
		ans = match && z.dp(i+1, j+1, s, p)
	}
	z.memo[k] = 1
	if ans {
		z.memo[k] = 2
	}
	return ans
}
