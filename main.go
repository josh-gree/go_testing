package main

import (
	"fmt"
)

func foo(x int) int {
	return 2 * x
}

func main() {
	x := 45
	y := foo(x)
	fmt.Println(y)
}