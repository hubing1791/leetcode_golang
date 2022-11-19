package medium_03_03_stack_of_plates_lcci

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	st := Constructor(2)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.PopAt(0))
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
}
