package easy_7_reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	type TestExample struct {
		num    int
		result int
	}

	testExamples := []TestExample{
		{123, 321},
		{456, 654},
	}

	for i, testExample := range testExamples {
		if testExample.result == reverse(testExample.num) {
			fmt.Printf("第%d个测试用例通过\n", i+1)
		}
	}
}
