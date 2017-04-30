package blah

import "testing"

func TestFoo(t *testing.T) {
	x := Foo(2)
	if x != 4 {
		t.Error("Expected 4 got ", x)
	}
}

func TestBar(t *testing.T) {
	x := Bar([]float64{1.1, 2.4})
	if x != 3.5 {
		t.Error("Expected 3.5 got ", x)
	}
}
