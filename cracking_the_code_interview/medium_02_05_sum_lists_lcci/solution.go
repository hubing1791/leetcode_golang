package medium_02_05_sum_lists_lcci

import (
	ll "leetcode_golang/custom_moudle/singly_linked_list"
)

//https://leetcode.cn/problems/sum-lists-lcci/
//2022-07-29

//这样写貌似每次都新建了节点，考虑写一个不需要新建的版本
func addTwoNumbers(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	digit, newHead := 0, &ll.ListNode{}
	newh := newHead
	left, right := 0, 0
	for l1 != nil || l2 != nil {
		left, right = 0, 0
		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}
		tmpNode := &ll.ListNode{(right + left + digit) % 10, nil}
		if right+left+digit > 9 {
			digit = 1
		} else {
			digit = 0
		}
		newh.Next = tmpNode
		newh = newh.Next

	}
	if digit == 1 {
		newh.Next = &ll.ListNode{1, nil}
	}
	return newHead.Next
}

//这样写貌似每次都新建了节点，考虑写一个不需要新建的版本
func addTwoNumbers1(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	digit, newHead := 0, &ll.ListNode{}
	newh := newHead
	left, right := 0, 0
	var tmpNode *ll.ListNode
	for l1 != nil || l2 != nil {
		left, right = 0, 0
		if l1 != nil {
			left = l1.Val
			tmpNode = l1
			l1 = l1.Next
		}
		if l2 != nil {
			right = l2.Val
			tmpNode = l2
			l2 = l2.Next
		}
		tmpNode.Val = (right + left + digit) % 10
		//一开始这儿还用了if else 属实没必要
		digit = (right + left + digit) / 10
		newh.Next = tmpNode
		newh = newh.Next

	}
	if digit == 1 {
		newh.Next = &ll.ListNode{1, nil}
	}
	return newHead.Next
}
