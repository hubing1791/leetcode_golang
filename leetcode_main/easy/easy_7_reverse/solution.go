package easy_7_reverse

import "math"

func reverse(x int) int {
	var result int = 0
	for x != 0 {
		if result < math.MinInt/10 || result > math.MaxInt/10 {
			return 0
		} else {
			tmpDigit := x % 10
			x /= 10
			result = tmpDigit + result*10
		}
	}
	return result
}
