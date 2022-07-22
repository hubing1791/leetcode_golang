package medium_179_largest_number

import (
	"sort"
	"strconv"
)

//2022-07-14
//https://leetcode.cn/problems/largest-number/

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		//之前这儿漏了等于，就出错了
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	var result_ []string
	for _, i := range nums {
		result_ = append(result_, strconv.Itoa(i))
	}
	result := ""
	for _, str := range result_ {
		result += str
	}
	return result
}
