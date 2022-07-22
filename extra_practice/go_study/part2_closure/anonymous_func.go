package part2_closure

import "fmt"

func anonFunc() {
	f := func() {
		fmt.Println("anonymous func without argument")
	}
	f()
	fmt.Printf("the addr of f() %X\n", &f)
	f()
	fmt.Printf("the addr of f() %X\n", &f)
}

func anonFuncWithArgs(args1 int, args2 int) {
	f := func(args1 int, args2 int) {
		fmt.Printf("anonymous func with two argument %d and %d\n", args1, args2)
	}
	f(args1, args2)
	fmt.Printf("the addr of f() %X\n", &f)
	f(args1+1, args2+1)
	fmt.Printf("the addr of f() %X\n", &f)

	func(args string) {
		fmt.Println(args)
	}("call the anonymous func immediately")
}

func anonFuncWithReturn() string {
	return func() string {
		return "return hhh"
	}()
}

func multiAnonFunc() {
	F := func(a int, b int) (func(int) int, func() int) {
		f1 := func(c int) int {
			return (a + b) * c
		}
		f2 := func() int {
			return a * b
		}
		return f1, f2
	}
	f1, f2 := F(1, 2)
	fmt.Printf("f1 return %d with args %d\nf2 return %d\n", f1(3), 3, f2())
}
