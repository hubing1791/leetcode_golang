package easy_1_two_sum

import (
	"sort"
)

// 找到的数字是对的，但是原题要求返回索引，所以不能先排序。
func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for {
		if start < end {
			if nums[start]+nums[end] == target {
				return []int{start, end}
			} else if nums[start]+nums[end] < target {
				start += 1
			} else {
				end -= 1
			}
		} else {
			break
		}
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	hashtable := make(map[int]int)
	for i, num := range nums {
		j, ok := hashtable[target-num]
		if !ok {
			hashtable[num] = i
		} else {
			return []int{i, j}
		}
	}
	return nil
}
