package infra

import (
	"io/ioutil"
	"net/http"
)

// Retriever 接收者
type Retriever struct{}

// Get 大写的Get， 作为public method
// 因为接收者里面啥也没有，所以不需要给接收者去一个名字
func (Retriever) Get(url string) string {
	r, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	bytes, _ := ioutil.ReadAll(r.Body)
	return string(bytes)
}
