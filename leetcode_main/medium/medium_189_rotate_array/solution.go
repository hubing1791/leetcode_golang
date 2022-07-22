package medium_189_rotate_array

//https://leetcode.cn/problems/rotate-array/
//2022-07-20

func rotateSlice(nums []int, k int) {
	length := len(nums)
	index := (length - k) % length
	//做这一步是因为，-9%2结果为-1，go无法索引负数
	if index < 0 {
		index += length
	}
	newNums := append(nums[index:], nums[:index]...)
	copy(nums, newNums)
}

// 这个版本的问题在于一旦k和length有倍数关系，就无法遍历数组
func rotateTwoPointer(nums []int, k int) {
	tmpNum := nums[0]
	length, count := len(nums), 0
	k = k % length
	pre, cur := length-k, 0
	//只需要循环换6次
	for ; count < length-1; count++ {
		nums[cur] = nums[pre]
		cur, pre = pre, pre-k
		if pre < 0 {
			pre += length
		}
	}
	nums[cur] = tmpNum
}

//上面一个方法的修正版本
func rotateTwoPointerRight(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

//找到最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func rotateReverse(nums []int, k int) {
	f := func(nums []int) {
		for i, n := 0, len(nums)/2; i < n; i++ {
			nums[i], nums[len(nums)-i] = nums[len(nums)-i], nums[i]
		}
	}
	k = k % len(nums)
	f(nums)
	f(nums[:k])
	f(nums[k:])
}
