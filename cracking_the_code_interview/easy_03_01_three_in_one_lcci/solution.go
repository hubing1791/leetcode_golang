package easy_03_01_three_in_one_lcci

//https://leetcode.cn/problems/three-in-one-lcci/
//2022-09-19

type TripleInOne struct {
	stack []int
	ptr   []int
	cap   int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack: make([]int, stackSize*3),
		ptr:   []int{-1, stackSize - 1, stackSize*2 - 1},
		cap:   stackSize,
	}
}

func (t *TripleInOne) Push(stackNum int, value int) {
	if t.ptr[stackNum] < t.cap*(stackNum+1)-1 {
		t.ptr[stackNum] += 1
		t.stack[t.ptr[stackNum]] = value
	}
}

func (t *TripleInOne) Pop(stackNum int) int {
	if t.ptr[stackNum] > t.cap*(stackNum)-1 {
		t.ptr[stackNum] -= 1
		return t.stack[t.ptr[stackNum]+1]
	}
	return -1
}

func (t *TripleInOne) Peek(stackNum int) int {
	if t.ptr[stackNum] > t.cap*(stackNum)-1 {
		return t.stack[t.ptr[stackNum]]
	}
	return -1
}

func (t *TripleInOne) IsEmpty(stackNum int) bool {
	return t.ptr[stackNum] == t.cap*(stackNum)-1
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
