package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// Get 实现者， 实现者不需要说他自己实现了哪个接口， 使用者规定了， 我这个必须有Get方法
// 这里这个类型 r *Retriever 是实现者要传给 接口的类型
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
