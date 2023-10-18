package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// func eval(a, b int, op string) (int, int)

// func eval (a, b int, op string) (q, r int) {
// 	q = 1
// 	r = 2
// 	return
// }

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		// panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling func %s with args "+"(%d %d) ", opName, a, b)
	return op(a, b)
}

// func eval(a, b int, op string) int {
// 	switch op {
// 	case "+":
// 		return a + b
// 	case "-":
// 		return a - b
// 	case "*":
// 		return a * b
// 	case "/":
// 		return a / b
// 	default:
// 		panic("unsupported operation: " + op)
// 	}
// }

func div(a, b int) (int, int) {
	return a / b, a % b
}

// func pow(a, b int) int {
// 	return int(math.Pow(float64(a), float64(b)))
// }

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	// fmt.Println(eval(3, 4, "x"))

	// 推荐使用这种方式处理
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	// fmt.Println(div(13, 3))
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
