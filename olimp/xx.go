package olimp

import "fmt"

func desigion(s0, m0, d int) {
	var balance, day, m int
	balance = s0
	for {
		balance += m * d
		balance -= m0
		m++
		day++
		fmt.Println(balance, day, m)
		if balance < 0 {
			break
		}
	}
}
