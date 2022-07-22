package medium_172_factorial_trailing_zeroes

// https://leetcode.cn/problems/factorial-trailing-zeroes/

func trailingZeroes(n int) int {
	result := 0
	for ; n > 0; n /= 5 {
		result += n / 5
	}
	return result
}
