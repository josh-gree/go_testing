package blah

import "testing"

type testpairFoo struct {
	input  int
	result int
}

type testpairBar struct {
	input  []float64
	result float64
}

var testsFoo = []testpairFoo{
	{2, 6},
	{0, 0},
	{-12, -36},
}

var testsBar = []testpairBar{
	{[]float64{1.1, 1.1}, 2.2},
	{[]float64{-1.1, 1.1}, 0},
	{[]float64{1.0, 2.0, 3.0, 0.0}, 6.0},
	{[]float64{}, 0.0},
}

func TestFoo(t *testing.T) {
	for _, pair := range testsFoo {
		v := Foo(pair.input)
		if v != pair.result {
			t.Error(
				"For", pair.input,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestBar(t *testing.T) {
	for _, pair := range testsBar {
		v := Bar(pair.input)
		if v != pair.result {
			t.Error(
				"For", pair.input,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
