package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		// 如果这里不传参的话， 那么就相当于是闭包， 外部的i就是自由变量， 在main这个goroutine结束之后
		// i的值为10， 此时再执行 a[i]++ 会 out of range
		// race condition i
		// 资源竞争
		// 可以使用channel来解决，暂时先不处理
		//i := i
		go func(i int) { // 主程序还在往下面跑， 只是并发地开了一个函数
			for {
				// 因为是io操作，所以会有一个等待的过程， 如果是操作一个数组的话，可能不会主动交出控制权
				fmt.Printf("Hello from "+"Goroutines-channel %d\n", i)
				//a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
