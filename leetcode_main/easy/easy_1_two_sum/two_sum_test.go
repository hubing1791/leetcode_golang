package __easy_two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T)  {
	type testExample struct{
		nums []int
		target int
	}
	testSet := []testExample{
		{[]int{1,2,3,4,5,6},6},
		{[]int{3,2,4},6},
	}
	for _,testExampleTemp := range testSet{
		fmt.Println(twoSum(testExampleTemp.nums,testExampleTemp.target))
	}
}