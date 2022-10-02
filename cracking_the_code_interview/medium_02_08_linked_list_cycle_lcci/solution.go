package medium_02_08_linked_list_cycle_lcci

//https://leetcode.cn/problems/linked-list-cycle-lcci/
//2022-09-19

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	nodeSet := make(map[*ListNode]struct{})
	newHead := head

	for newHead != nil {
		_, ok := nodeSet[newHead]
		if ok {
			return newHead

		} //else {
		//	nodeSet[newHead] = struct{}{}
		//}
		// it is unnecessary to use else
		nodeSet[newHead] = struct{}{}
		newHead = newHead.Next
	}
	return nil
}
