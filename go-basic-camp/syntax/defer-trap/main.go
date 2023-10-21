package main

// 从地址的角度思考即可！
// V1: 2，2，2， 用的是同一个地址
// V2: 2，1，0， 用的是不同的地址
// V3: 2，1，0， 用的是不同的地址

func main() {
	//deferClosureLoopV1()
	//deferClosureLoopV2()
	deferClosureLoopV3()
}
