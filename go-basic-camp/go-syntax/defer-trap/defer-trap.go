package main

func deferClosureLoopV1() {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}
}

func deferClosureLoopV2() {
	for i := 0; i < 3; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func deferClosureLoopV3() {
	for i := 0; i < 3; i++ {
		val := i
		defer func() {
			println(val)
		}()
	}
}
