package main

import (
	"fmt"
	"time"

	"github.com/upupnoah/go-basic-future/go-basic-godev/retriever/mock"
	"github.com/upupnoah/go-basic-future/go-basic-godev/retriever/real"
)

// Retriever 定义了一个接口，实现者需要有 Get方法
type Retriever interface {
	Get(url string) string
}

// 直接通过接口调用Get方法
func download(r Retriever) string {
	return r.Get("https://www.ascwing.com")
}

func main() {
	var r Retriever = &mock.Retriever{Contents: "This is a fake url"}
	inspect(r)
	//fmt.Println(download(r))
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever!")
	}
	fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
