package reverse_integer

import "math"

const (
	maxi = (1<<31 - 1) / 10
	mini = (-1 << 31) / 10
)

func reverse(x int) int {
	var i int32
	xx := int32(x)
	for xx != 0 {
		dx := xx % 10
		if i > maxi {
			return 0
		}
		if i < mini {
			return 0
		}
		i *= 10
		if dx > 0 && i > 0 && math.MaxInt32-dx < i {
			return 0
		}
		if dx < 0 && i < 0 && math.MinInt32-dx > i {
			return 0
		}
		i += dx
		xx -= dx
		xx /= 10
	}
	return int(i)
}
