package part2_closure

import (
	"fmt"
	"testing"
)

func Test_anonFunc(t *testing.T) {
	anonFunc()
}

func Test_anonFuncWithArgs(t *testing.T) {
	anonFuncWithArgs(1, 2)
}

func Test_anonFuncWithReturn(t *testing.T) {
	fmt.Println(anonFuncWithReturn())
}

func Test_multiAnonFunc(t *testing.T) {
	multiAnonFunc()
}
