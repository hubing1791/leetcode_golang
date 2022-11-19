package medium_03_03_stack_of_plates_lcci

//https://leetcode.cn/problems/stack-of-plates-lcci/
//2022-09-19
type StackOfPlates struct {
	stack [][]int
	cap   int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		stack: make([][]int, 0),
		cap:   cap,
	}
}

func (s *StackOfPlates) Push(val int) {
	numStack := len(s.stack)
	var lengthLast int
	if numStack != 0 {
		lengthLast = len(s.stack[numStack-1])
	} else {
		lengthLast = s.cap
	}
	// add a new stack or push in current stack
	if lengthLast == s.cap {
		s.stack = append(s.stack, []int{val})
	} else {
		s.stack[numStack-1] = append(s.stack[numStack-1], val)
	}
}

func (s *StackOfPlates) Pop() int {
	numStack := len(s.stack)
	if numStack == 0 {
		return -1
	}
	lengthLast := len(s.stack[numStack-1])
	result := s.stack[numStack-1][lengthLast-1]
	if lengthLast == 1 {
		s.stack = s.stack[:numStack-1]
	} else {
		s.stack[numStack-1] = s.stack[numStack][:lengthLast-1]
	}
	return result
}

func (s *StackOfPlates) PopAt(index int) int {
	numStack := len(s.stack)
	if index+1 > numStack {
		return -1
	}
	lengthIndex := len(s.stack[index])
	result := s.stack[index][lengthIndex-1]
	if lengthIndex == 1 {
		s.stack = append(s.stack[:index], s.stack[index+1:]...)
	} else {
		s.stack[index] = s.stack[index][:lengthIndex-1]
	}
	return result
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
