package main

import "fmt"

func main() {
	for _, v := range "SB" {
		fmt.Printf("%T %v %c\n   ", v, v, v)
	}
}
