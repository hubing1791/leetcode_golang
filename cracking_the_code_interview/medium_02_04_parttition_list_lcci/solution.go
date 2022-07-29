package medium_02_04_parttition_list_lcci

import (
	ll "leetcode_golang/custom_moudle/singly_linked_list"
)

//https://leetcode.cn/problems/partition-list-lcci/
//2022-07-25

func partition(head *ll.ListNode, x int) *ll.ListNode {
	newHeadSmall, newHeadBig := &ll.ListNode{}, &ll.ListNode{}
	small, big := newHeadSmall, newHeadBig
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	//之前进行了small.next = nil没必要
	big.Next = nil
	// 之前加了个毫无意义的if newSmallHead.Next != nil的条件，没啥意义
	small.Next = newHeadBig.Next

	return newHeadSmall.Next
}
