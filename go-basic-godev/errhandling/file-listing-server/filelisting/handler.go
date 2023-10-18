package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/NoahX/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//fmt.Println(request.URL.Path)
	//fmt.Println(prefix)
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start " + "with " + prefix)
	}
	// 去掉路径前缀，只留下 filename
	path := request.URL.Path[len("/NoahX/"):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(all)
	return nil
}
