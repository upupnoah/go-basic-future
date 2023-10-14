package main

import "fmt"

func main() {
	// ******* map 变量的声明和初始化 *******

	// 不赋初值的 map 不可用，用直接 panic
	//var m map[string]int // m = nil
	//m["key"] = 1         // 发生运行时异常：panic: assignment to entry in nil map

	// 赋初值的 map
	m := map[string]int{}
	m["key"] = 1

	// 1. 使用复合字面值初始化 map 类型变量
	_ = map[int][]string{
		1: []string{"val1_1", "val1_2"},
		3: []string{"val3_1", "val3_2", "val3_3"},
		7: {"val7_1"}, // Go 允许省略字面值中的元素类型
	}

	type Position struct {
		x float64
		y float64
	}

	// 语法糖，可以省略 Position
	_ = map[Position]string{
		Position{29.935523, 52.568915}: "school",
		{25.352594, 113.304361}:        "shopping-mall",
		{73.224455, 111.804306}:        "hospital",
	}

	// 2. 使用 make 创建 map 类型变量
	_ = make(map[int]string)    // 未指定初始容量
	_ = make(map[int]string, 8) // 指定初始容量为8

	// ******* map 的基本操作 *******

	// 1. 插入新键值对
	m = make(map[string]int)
	m["key1"] = 1

	// 2. 获取键值对数量
	fmt.Println(len(m)) // 1
	m["key3"] = 3
	fmt.Println(len(m)) // 2

	// 3. 查找和数据读取
	m = make(map[string]int)
	_ = m["key1"] // 这样无法判断 key 是否存在，如果不存在，返回的是 value 类型的零值（需要通过 comma ok 判断）

	// comma ok: 判断 map 中是否存在某个 key
	// 在 Go 语言中，请使用“comma ok”惯用法对 map 进行键查找和键值读取操作!!!!
	m = make(map[string]int)
	m["key1"] = 1
	if v, ok := m["key1"]; ok {
		fmt.Printf("key1 is in map, value is %d\n", v)
	} else {
		fmt.Println("key1 is not in map")
	}

	// 4. 删除数据
	// delete 函数是从 map 中删除键的唯一方法
	// 即便传给 delete 的键在 map 中并不存在，delete 函数的执行也不会失败，更不会抛出运行时的异常
	delete(m, "key1")

	// 5. 遍历 map 中的键值数据
	m1 := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	fmt.Printf("{ ")
	for k, v := range m1 {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Println()

	//for k, _ := range m1 { // 只需要 key
	//	// 使用k
	//	fmt.Println(k)
	//}
	//for k := range m1 { // 只需要 key，更地道的写法
	//	// 使用k
	//	fmt.Println(k)
	//}
	//for _, v := range m1 { // 只关心 v
	//	// 使用v
	//	fmt.Println(v)
	//}

	// 对同一 map 做多次遍历的时候，每次遍历元素的次序都不相同
	// 程序逻辑千万不要依赖遍历 map 所得到的的元素次序
}
