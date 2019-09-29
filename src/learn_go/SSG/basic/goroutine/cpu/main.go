package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())

	//自己设置使用多少个CPU
	runtime.GOMAXPROCS(4)
}
