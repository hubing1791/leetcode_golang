package singly_linked_list

import "testing"

func TestLinked(T *testing.T){
	head := Generator([]int{1,2,3,4,5})
	ShowLL(head)
}
