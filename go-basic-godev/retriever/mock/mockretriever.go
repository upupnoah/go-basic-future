package mock

import "fmt"

// Retriever 结构体
type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	//panic("implement me")
	return fmt.Sprintf("Retriever: {Contents = %s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// Get 为结构体实现方法
//     接收者
func (r *Retriever) Get(url string) string {
	return r.Contents
}
