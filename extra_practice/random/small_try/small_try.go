package main

import (
	"fmt"
)

func mapChange(m map[int]int){
	m[1]=5
	m[5]=1
}

func main() {
	//不同的遍历字符串的方式，返回值不同
	//strTry := "abcdefg"
	//fmt.Printf("%+v,		%v\n", strTry[0],reflect.TypeOf(strTry[0]))
	//for _,char :=range strTry {
	//	fmt.Printf("%+v,		%v\n",char,reflect.TypeOf(char))
	//}
	//

	var mapTry map[int]int
	mapTry = make(map[int]int)
	mapChange(mapTry)
	fmt.Println(mapTry)
}
