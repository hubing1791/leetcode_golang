package heap_sort

import (
	"container/heap"
	"fmt"
	"testing"
)

func generateTestData(intList []int) []myStruct {
	result := make([]myStruct, len(intList))
	for i := 0; i < len(intList); i++ {
		result[i] = myStruct{name: "A", ForS: sortStruct{intList[i], true}}
	}
	return result
}

func TestComplexHeap(t *testing.T) {
	testSample := []int{7, 9, 4, 8, 5, 6, 1}
	myHeapT := &myHeap{}
	heap.Init(myHeapT)
	for _, st := range generateTestData(testSample) {
		heap.Push(myHeapT, st)
	}
	for i := 0; i < len(testSample); i++ {
		fmt.Println(heap.Pop(myHeapT))
	}
}
