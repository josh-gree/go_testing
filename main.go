package main

import (
	"fmt"
)

func foo(x int) int {
	return 2 * x
}

func main() {
	z := []int{12, 34, 56, 45, 34, 12, 32, 13}
	for _, x := range z {
		fmt.Println(foo(x))
	}
}
