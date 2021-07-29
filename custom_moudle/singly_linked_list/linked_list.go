package singly_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Generator 根据给的数组切片生成单链表
func Generator(nums []int) *ListNode {
	var head = ListNode{0, nil}
	pointer := &head
	for _, num := range nums {
		var tempNode = ListNode{num, nil}
		pointer.Next = &tempNode
		pointer = &tempNode
	}
	return head.Next
}

func ShowLL(node *ListNode){
	var result []int
	for node != nil{
		result = append(result,node.Val)
		node=node.Next
	}
	fmt.Println(result)
}