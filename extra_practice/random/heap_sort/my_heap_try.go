package heap_sort

type sortStruct struct {
	ForSort int  // sort by this value
	sth     bool // simulate a complex structure
}
type myStruct struct {
	name string // simulate a complex structure
	ForS sortStruct
}

type myHeap []myStruct

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i, j int) bool {
	return m[i].ForS.ForSort < m[j].ForS.ForSort
}

func (m myHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func (m *myHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*m = append(*m, x.(myStruct))
}
