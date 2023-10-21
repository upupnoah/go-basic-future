package main

func main() {
	var x int32 = 1
	x = ^(x << 31)
	println(x)
	println(x+1, x+2, x+10)
}

type T1[T int | string | ~float64] struct {
}

type myInterface interface {
	comparable
	int | string | ~float64
}

func maxstring[T string](a, b T) T {
	if a > b {
		return a
	}
	return b
}
