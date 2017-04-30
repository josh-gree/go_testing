package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func foo(x int) int {
	return 2 * x
}

func main() {
	fmt.Println("Hello")

	e := echo.New()
	e.Start(":9000")
}
