package main

import "fmt"

/*
	map:
		init: make or {}
		op: delete() -> 删除元素
		for range order: disorderly

		recommend: if v, ok := m[key]; ok {...}
*/

func main() {
	// map[K]V, map[K1]map[K2]V
	// 这是一个hash map， 内部无序， 所以每次输出的顺序都是不一样的
	m := map[string]string{
		"name":    "NoahX",
		"course":  "Go-Development-Engineer",
		"site":    "hangzhou",
		"quality": "good",
	}

	m2 := make(map[string]int) // m2 == empty map
	fmt.Println(m, m2)

	var m3 map[string]int // m3 == nil
	fmt.Println(m3)

	fmt.Println("Traversing map...")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values...")
	courseName := m["course"] // 如何不存在，是一个Zero Value 空行
	fmt.Println(courseName)

	//优化
	fmt.Println("优化...")
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist...")
	}

	// 删除元素
	fmt.Println("Deleting values...")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(ok)
}
