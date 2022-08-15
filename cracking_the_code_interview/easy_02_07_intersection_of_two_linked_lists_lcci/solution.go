package easy_02_07_intersection_of_two_linked_lists_lcci

//https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/
//2022-08-10

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	newHeadA, newHeadB := headA, headB
	flagA, flagB := true, true
	//一开始没考虑有一个是空的情况
	if headA == nil || headB == nil {
		return nil
	}

	for newHeadA != nil || newHeadB != nil {
		if newHeadA == nil {
			if flagA {
				flagA = false
				newHeadA = headB
			} else {
				return nil
			}
		}
		if newHeadB == nil {
			if flagB {
				flagB = false
				newHeadB = headA
			} else {
				return nil
			}
		}
		//这个写在前面就不对
		if newHeadA == newHeadB {
			return newHeadA
		}
		newHeadA, newHeadB = newHeadA.Next, newHeadB.Next
	}
	return nil
}
