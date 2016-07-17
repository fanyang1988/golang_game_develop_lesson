package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	1.使用flag和os.Args
	flag包可以很方便的解析类似于“--v 100”这样的
	参数, 同时可以指定解析的类型
	通过os.Args 也可以获得完整的执行命令
*/

// some params from flags
var (
	// paramV a int flags
	paramV int
	// paramS a string flags
	paramS string
)

func useFlag() {
	// os.Args
	fmt.Printf("os Args %v\n", os.Args)

	// 使用flag解析 --v {param} 这样的参数
	flag.IntVar(&paramV, "v", -1, "a int")
	flag.StringVar(&paramS, "s", "", "a string")
	flag.Parse()

	fmt.Printf("param V %d param S %s\n", paramV, paramS)
}
