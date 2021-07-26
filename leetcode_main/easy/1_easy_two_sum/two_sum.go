package __easy_two_sum

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for {
		if start < end {
			if nums[start]+nums[end] == target {
				return []int{nums[start], nums[end]}
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
