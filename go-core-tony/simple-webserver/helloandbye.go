package main

import (
	"fmt"
	"io"
)

// ReadAll 通过接受接口类型参数的普通函数进行组合
// ReadAll函数通过io.Reader 这个接口， 完成了将io.Reader的实现与ReadAll所在的包低耦合地组合在一起
// 从而达到从任何 实现了 io.Reader的数据源中 读取数据的目的
func ReadAll(r io.Reader) ([]byte, error) {
	return nil, nil
}

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return 0, nil
}

func main() {
	var s1 []int
	//var s2 = []int{}
	s1 = append(s1, 1)
	fmt.Println(s1)
	//var str string

}
