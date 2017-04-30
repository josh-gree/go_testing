package blah

import (
	"fmt"
	"testing"
)

var tests = []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

func TestFoo(t *testing.T) {
	fn := Foo(23)
	fmt.Println(fn)

}
