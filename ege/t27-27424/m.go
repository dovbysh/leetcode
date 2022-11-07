package t27_27424

import "math"

func strageSum(a [][2]int) (int, error) {
	sum := 0
	for _, v := range a {
		if v[0] > v[1] {
			sum += v[0]
		} else {
			sum += v[1]
		}
	}
	if sum%3 == 0 {
		minDelta := math.MinInt64
		var d int
		for _, v := range a {
			if v[0] > v[1] {
				d = v[1] - v[0]
				if d > minDelta && d*(-1)%3 > 0 {
					minDelta = d
				}
			} else {
				d = v[0] - v[1]
				if d > minDelta && d*(-1)%3 > 0 {
					minDelta = d
				}
			}
		}
		sum += minDelta
	}
	return sum, nil
}
