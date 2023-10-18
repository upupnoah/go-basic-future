package main

import (
	"errors"
	"fmt"
)

// ErrSentinel 哨兵错误，将字符串作为错误容易出错（写代码的时候），因此将整个字符串定义为一个常量
var ErrSentinel = errors.New("the underlying sentinel error")

type MyError struct {
	e string
}

func (myErr *MyError) Error() string {
	return myErr.e
}

func main() {
	/*
		两个错误是不相等的，但他们是同一种类型的错误
	*/
	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == ErrSentinel) //false

	// errors.Is(上层的err， 底层的err）
	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
	} else {
		println("err2 is not ErrSentinel")
	}

	println("***********************")

	var err = &MyError{"MyError error demo"}
	err3 := fmt.Errorf("wrap err: %w", err)
	err4 := fmt.Errorf("wrap err3: %w", err3)
	var e *MyError
	if errors.As(err4, &e) {
		println("MyError is on the chain of err4")
		println(err == e)
		return
	}
	println("MyError is not on the chain of err4")
}
