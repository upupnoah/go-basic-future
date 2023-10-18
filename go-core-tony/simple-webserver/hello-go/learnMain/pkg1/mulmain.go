package pkg1

import "fmt"

var cnt int

type helloStruct struct {
	name string
	age  int
}

func (h helloStruct) getName() string {
	return h.name
}

func Main() {
	if cnt == 6 {
		return
	}
	cnt++
	main()
}

func main() {
	fmt.Println("main func for pkg1")
	//Main()
}
