package main

import (
	"fmt"
	"time"

	"github.com/upupnoah/go-basic-future/go-basic-godev/retriever/mock"
	"github.com/upupnoah/go-basic-future/go-basic-godev/retriever/real"
)

const url = "https://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func main() {
	var r Retriever
	retriever := mock.Retriever{
		Contents: "this is a fake QAQ.com",
	}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	fmt.Println("Try a session: ")
	fmt.Println(session(&retriever))
	fmt.Println()

	var toString fmt.Stringer
	toString = &mock.Retriever{Contents: "Hello QWQ..."}
	fmt.Println("ToString: ")
	fmt.Println(toString.String())
}

func session(s RetrieverPoster) string {
	//s.Get(url)
	//s.Post(url,
	//	map[string]string{
	//		"name":   "QAQ",
	//		"course": "Go-Development-Engineer",
	//	})
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
