package part2_closure

import "fmt"

//
// 内层的函数可以使用外层函数的所有变量，即便外层函数已经执行完了

func Fun() func(args string) string {
	a := "py "
	return func(args string) string {
		a += " " + args
		return a
	}
}

func ClosureTry1() {
	funcA := Fun()
	bStr := funcA("try")
	cStr := funcA("try")
	fmt.Println(bStr, "\n", cStr)
}

// ClosureTry2 两次返回的匿名函数地址是不同的
func ClosureTry2() {
	funcA := Fun()
	funcB := Fun()
	fmt.Printf("%X\n%X\n", &funcA, &funcB)
}

// ClosureTry3 模拟闭包引用的一种错误,因为只调用了一次funcSlice的生成函数，只形成了一个闭包
func ClosureTry3() {
	funcSlice := func() []func() {
		a := make([]func(), 3)
		for i := 0; i < 3; i++ {
			a[i] = func() {
				fmt.Println(&i, i)
			}
		}
		return a
	}()
	funcSlice[0]()
	funcSlice[1]()
	funcSlice[2]()
}

// ClosureTry4 3的纠正版本1,这种是将(i)先防御是复制并传到了匿名函数中
func ClosureTry4() {
	funcSlice := func() []func() {
		b := make([]func(), 3, 3)
		for i := 0; i < 3; i++ {
			b[i] = (func(j int) func() {
				return func() {
					fmt.Println(&j, j)
				}
			})(i)
		}
		return b
	}()
	funcSlice[0]()
	funcSlice[1]()
	funcSlice[2]()
}

// ClosureTry5 3的纠正版本2,这种是将(i)先防御是复制并传到了匿名函数中
func ClosureTry5() {
	funcSlice := func() []func() {
		b := make([]func(), 3, 3)
		for i := 0; i < 3; i++ {
			j := i
			b[i] = func() {
				fmt.Println(&j, j)
			}
		}
		return b
	}()
	funcSlice[0]()
	funcSlice[1]()
	funcSlice[2]()
}

// ClosureTryWithDefer 先执行r=1,再++
func ClosureTryWithDefer() {
	F := func() (r int) {
		defer func() {
			r++
		}()
		return 1
	}
	fmt.Println(F())
}
