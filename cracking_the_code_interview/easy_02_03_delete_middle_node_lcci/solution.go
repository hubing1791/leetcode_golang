package easy_02_03_delete_middle_node_lcci

import ll "leetcode_golang/custom_moudle/singly_linked_list"

//https://leetcode.cn/problems/delete-middle-node-lcci
//2022-07-25

func deleteNode(node *ll.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
