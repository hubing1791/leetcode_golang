package easy_02_02_kth_node_from_end_of_list_lcci

//https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
//2022-07-25
import ll "leetcode_golang/custom_moudle/singly_linked_list"

func kthToLast(head *ll.ListNode, k int) int {
	cyclePoi := head
	for i := 0; i < k; i++ {
		cyclePoi = cyclePoi.Next
	}
	resultNode := head
	for cyclePoi != nil {
		resultNode, cyclePoi = resultNode.Next, cyclePoi.Next
	}
	return resultNode.Val
}
