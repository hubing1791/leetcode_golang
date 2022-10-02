package medium_03_03_stack_of_plates_lcci

//https://leetcode.cn/problems/stack-of-plates-lcci/
//2022-09-19

type StackOfPlates struct {
	stack [][]int
	cap   int
	Idx   int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		stack: make([][]int, 0),
		cap:   cap,
		Idx:   -1,
	}
}

func (s *StackOfPlates) Push(val int) {
}

func (s *StackOfPlates) Pop() int {

}

func (s *StackOfPlates) PopAt(index int) int {

}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
