package medium_2_add_two_numbers

import (
	"leetcode_golang/custom_moudle/singly_linked_list"
	"testing"
)

func TestAdd(T *testing.T) {
	l1 := singly_linked_list.Generator([]int{1, 2, 3})
	l2 := singly_linked_list.Generator([]int{9, 4, 4})
	l3 := addTwoNumbersSimple(l1, l2)
	singly_linked_list.ShowLL(l3)

}
