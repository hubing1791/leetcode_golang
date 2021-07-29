package main

import (
	"fmt"
	"reflect"
)

func main() {
	strTry := "abcdefg"
	fmt.Printf("%+v,		%v\n", strTry[0],reflect.TypeOf(strTry[0]))
	for _,char :=range strTry {
		fmt.Printf("%+v,		%v\n",char,reflect.TypeOf(char))
	}
}
