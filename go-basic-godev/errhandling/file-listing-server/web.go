package main

import (
	"net/http"
	"os"

	"github.com/gpmgo/gopm/log"

	"github.com/upupnoah/go-basic-future/go-basic-godev/errhandling/file-listing-server/filelisting"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error
type handler func(http.ResponseWriter, *http.Request)

func errWrapper(handler appHandler) handler {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Warn("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())

			// UserError
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				//http.Error(
				//	writer,                               // 向谁 汇报error
				//	http.StatusText(http.StatusNotFound), // err.Error() 也可以， 但是不想暴露内部错误
				//	http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden

			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				writer,
				http.StatusText(code),
				code)
		}
	}
}

// 有些错误希望给用户看到， 有些则不希望给用户看到
type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
