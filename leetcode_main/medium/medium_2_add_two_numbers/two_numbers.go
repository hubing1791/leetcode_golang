package __medium_add_two_numbers

import "leetcode_golang/custom_moudle/singly_linked_list"

type ListNode  = singly_linked_list.ListNode
//顺便也可以用这个函数测试*传参直接改变参数的情况如何实现
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//进位
	carry := 0
	var newTail = ListNode{0, nil}
	var pointer = &newTail
	var tempNum int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tempNum = l2.Val + carry
			if tempNum >= 10 {
				tempNum -= 10
				carry = 1
				var tempNode = ListNode{tempNum, nil}
				pointer.Next = &tempNode
				pointer = &tempNode
			} else {
				var tempNode = ListNode{tempNum, nil}
				pointer.Next = &tempNode
				pointer = &tempNode
				carry = 0
			}
			l2 = l2.Next
		} else if l2 == nil {
			tempNum = l1.Val + carry
			if tempNum >= 10 {
				tempNum -= 10
				carry = 1
				var tempNode = ListNode{tempNum, nil}
				pointer.Next = &tempNode
				pointer = &tempNode
			} else {
				var tempNode = ListNode{tempNum, nil}
				pointer.Next = &tempNode
				pointer = &tempNode
				carry = 0
			}
			l1 = l1.Next
		} else {
			tempNum = l1.Val + l2.Val+carry
			if tempNum >= 10 {
				tempNum -= 10
				carry = 1
				var tempNode = ListNode{tempNum, nil}
				pointer.Next = &tempNode
				pointer = &tempNode
			} else {
				var tempNode = ListNode{tempNum, nil}
				pointer.Next = &tempNode
				pointer = &tempNode
				carry = 0
			}
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if carry == 1{
		var tempNode = ListNode{1, nil}
		pointer.Next = &tempNode
	}
	return newTail.Next
}

func addHelper(pointer *ListNode, num int, carry *int)*ListNode {
	if num>=10{
		num -= 10
		*carry = 1
		var tempNode = ListNode{num, nil}
		pointer.Next = &tempNode
		return &tempNode
	}else {
		*carry = 0
		var tempNode = ListNode{num, nil}
		pointer.Next = &tempNode
		return &tempNode
	}

}

//把原版的分解为两个函数以写的精简一些
func addTwoNumbersSimple(l1 *ListNode, l2 *ListNode) *ListNode{
	carry := 0
	var newTail = ListNode{0, nil}
	var pointer = &newTail
	var tempNum int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tempNum = l2.Val + carry
			pointer = addHelper(pointer,tempNum,&carry)
			l2 = l2.Next
		} else if l2 == nil {
			tempNum = l1.Val + carry
			pointer = addHelper(pointer,tempNum,&carry)
			l1 = l1.Next
		} else {
			tempNum = l1.Val + l2.Val+carry
			pointer = addHelper(pointer,tempNum,&carry)
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if carry == 1{
		var tempNode = ListNode{1, nil}
		pointer.Next = &tempNode
	}
	return newTail.Next
}