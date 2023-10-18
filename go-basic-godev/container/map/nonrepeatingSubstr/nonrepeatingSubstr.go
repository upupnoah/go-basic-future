// 寻找最长不含有重复字符的子串
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfNonRepeatingSubstr(s string) int {
	ans, j := 0, 0
	str := []rune(s)
	m := map[rune]int{}
	for i, v := range str {
		m[v]++
		for m[v] > 1 {
			m[str[j]]--
			j++
		}
		ans = max(ans, i-j+1)
	}
	return ans
}

func main() {
	// 寻找最长不包含重复字符的子串
	var s string
	fmt.Scan(&s)
	// str := []rune(s) // 转成rune即可支持中文
	// for i, j := 0, 0; i < len(str); i++ {
	// 	m[str[i]]++
	// 	for m[str[i]] > 1 {
	// 		m[str[j]]--
	// 		j++
	// 	}
	// 	ans = max(ans, i-j+1)
	// }
	fmt.Println(lengthOfNonRepeatingSubstr(s))
}
