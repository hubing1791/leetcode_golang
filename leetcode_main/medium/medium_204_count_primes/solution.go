package medium_204_count_primes

//https://leetcode.cn/problems/count-primes/
//2022-07-20

//想到是埃式筛法，go还是有点不熟
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	var result = 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			result += 1
			for j := i * 2; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return result
}
