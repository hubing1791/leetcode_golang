package main

import (
	"fmt"
	"reflect"
)

//实验一下map查找不到时的返回值

type People struct {
	name string
	age  int
}

func main() {
	//value不同时，看看map初始值的区别
	var strAndInt = make(map[string]int)
	var strAndEmptyStruct = make(map[string]struct{})
	var strAndStruct = make(map[string]People)
	var strAndStr = make(map[string]string)
	fmt.Printf("value是int时，找不到key不存在时对应的值为%+v,类型为%v\n", strAndInt["hhh"], reflect.TypeOf(strAndInt["hhh"]))
	fmt.Printf("value是空结构体时，找不到key不存在时对应的值为%+v,类型为%v\n", strAndEmptyStruct["hhh"], reflect.TypeOf(strAndEmptyStruct["hhh"]))
	fmt.Printf("value是非空结构体时，找不到key不存在时对应的值为%+v,类型为%v\n", strAndStruct["hhh"], reflect.TypeOf(strAndStruct["hhh"]))
	fmt.Printf("value是字符串时，找不到key不存在时对应的值为%+v,类型为%v\n", strAndStr["hhh"], reflect.TypeOf(strAndStr["hhh"]))
	if _,ok := strAndStruct["hh"];ok{
		fmt.Println("找不到也ok")
	}
	if _,ok := strAndStruct["hh"];ok{
		fmt.Println("再找就ok")
	}
	strAndStruct["hh"] = People{}
	if _,ok := strAndStruct["hh"];ok{
		fmt.Println("用空结构初始化就ok")
	}
	tempP := strAndStruct["hhhhhhhhhhh"]
	fmt.Println(tempP)
}
