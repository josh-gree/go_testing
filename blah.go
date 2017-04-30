package blah

// Foo multiplies by 2
func Foo(x int) int {
	return 3 * x
}

// Bar returns the sum of a slice of floats
func Bar(x []float64) float64 {
	out := 0.0
	for _, num := range x {
		out += num
	}
	return out
}
