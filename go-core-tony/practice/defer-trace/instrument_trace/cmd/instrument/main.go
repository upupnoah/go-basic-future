package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/upupnoah/go-basic-future/go-core-tony/practice/defer-trace/instrument_trace/instrumenter"
	"github.com/upupnoah/go-basic-future/go-core-tony/practice/defer-trace/instrument_trace/instrumenter/ast"
)

var (
	wrote bool
)

func init() {
	flag.BoolVar(&wrote, "w", false, "write result to (source) file instead of stdout")
}

func usage() {
	fmt.Println("instrument [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	flag.Parse() // 解析命令行参数

	if len(os.Args) < 2 { // 对命令行参数个数进行校验
		usage()
		return
	}

	var file string
	if len(os.Args) == 3 {
		file = os.Args[2]
	}

	if len(os.Args) == 2 {
		file = os.Args[1]
	}
	if filepath.Ext(file) != ".go" { // 对源文件扩展名进行校验
		usage()
		return
	}

	var ins instrumenter.Instrumenter // 声明instrumenter.Instrumenter接口类型变量

	// 创建以ast方式实现Instrumenter接口的ast.instrumenter实例
	ins = ast.New("github.com/upupnoah/go-basic-godev-future/go-core-tony/practice/defer-trace/instrument_trace", "trace", "Trace")
	newSrc, err := ins.Instrument(file) // 向Go源文件所有函数注入Trace函数
	if err != nil {
		panic(err)
	}

	if newSrc == nil {
		// add nothing to the source file. no change
		fmt.Printf("no trace added for %s\n", file)
		return
	}

	if !wrote {
		fmt.Println(string(newSrc)) // 将生成的新代码内容输出到stdout上
		return
	}

	// 将生成的新代码内容写回原Go源文件
	if err = os.WriteFile(file, newSrc, 0666); err != nil {
		fmt.Printf("write %s error: %v\n", file, err)
		return
	}
	fmt.Printf("instrument trace for %s ok\n", file)
}
