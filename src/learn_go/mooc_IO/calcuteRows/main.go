package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	//通过Edit Configurations中修改Program arguments
	//os.Args	//启动时，从命令行传进来的参数，第一个参数是要运行的程序
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line int
	for{
		_,isPrefix,err := reader.ReadLine()	//如果一行过长，则isPrefix = true
		if err != nil {
			break
		}
		if !isPrefix{
			line ++
		}
	}
	fmt.Println(line)
}
