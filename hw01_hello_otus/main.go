package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

var msg="Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(msg))
}