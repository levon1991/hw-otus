package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

var msg = "Hello, OTUS!"

func test() bool {
	return true
}
func main() {
	fmt.Println(stringutil.Reverse(msg))
	if test() {
		fmt.Println("1")
	} else {
		if !test() {
			fmt.Println("2")
		} else if test() == true {
			fmt.Println("3")
		} else {
			fmt.Println("4")
		}
	}
}
