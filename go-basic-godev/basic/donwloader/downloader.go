package main

import (
	"fmt"

	"Learn_go/Go-Development-Engineer/basic/donwloader/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

// ?: Something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	// 为了低耦合
	url := "https://www.baidu.com"
	// retriever := infra.Retriever{}
	// var retriever testing.Retriever = getRetriever()
	var r retriever = getRetriever()
	// retriever := getRetriever()
	fmt.Println(r.Get(url))
}
