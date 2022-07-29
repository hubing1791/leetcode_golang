package easy_02_06_palindrome_linked_list_lcci

import (
	ll "leetcode_golang/custom_moudle/singly_linked_list"
)

//https://leetcode.cn/problems/palindrome-linked-list-lcci/
//2022-07-29

// 使用递归的方法
func isPalindrome(head *ll.ListNode) bool {
	frontPointer := head
	var reCheck func(node *ll.ListNode) bool
	reCheck = func(node *ll.ListNode) bool {
		if node != nil {
			if !reCheck(node.Next) {
				return false
			}
			if node.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return reCheck(head)
}

//符合进阶需求的

func reLL(head *ll.ListNode) *ll.ListNode {
	cur := head
	var pre *ll.ListNode
	pre = nil
	for cur != nil {
		cur, pre, cur.Next = cur.Next, cur, pre
	}
	return pre
}

func findHalf(head *ll.ListNode) *ll.ListNode {
	fast := head
	for fast != nil {
		head = head.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return head
}
func isPalindrome1(head *ll.ListNode) bool {
	half := findHalf(head)
	halfRe := reLL(half)
	for head != nil && halfRe != nil {
		if head.Val != halfRe.Val {
			return false
		}
		head = head.Next
		halfRe = halfRe.Next
	}
	return true
}
