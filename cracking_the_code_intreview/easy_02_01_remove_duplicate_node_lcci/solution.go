package easy_02_01_remove_duplicate_node_lcci

//https://leetcode.cn/problems/remove-duplicate-node-lcci/
//2022-07-22
import ll "leetcode_golang/custom_moudle/singly_linked_list"

func removeDuplicateNodes1(head *ll.ListNode) *ll.ListNode {
	if head == nil {
		return head
	}
	valSet := map[int]bool{head.Val: true}
	cycleHead := head
	for cycleHead.Next != nil {
		if !valSet[cycleHead.Next.Val] {
			valSet[cycleHead.Next.Val] = true
			cycleHead = cycleHead.Next
		} else {
			//一开始这儿忘记写一个Next了
			cycleHead.Next = cycleHead.Next.Next
		}
	}
	cycleHead.Next = nil
	return head
}

func removeDuplicateNodes(head *ll.ListNode) *ll.ListNode {
	cycle1 := head
	for cycle1 != nil {
		cycle2 := cycle1
		for cycle2.Next != nil {
			if cycle2.Next.Val == cycle1.Val {
				cycle2.Next = cycle2.Next.Next
			} else {
				cycle2 = cycle2.Next
			}
		}
		cycle1 = cycle1.Next
	}
	return head
}
