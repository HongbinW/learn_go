package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader,num int)([] byte,error){//字符串读取器，只要来源实现了io.Reader接口

	//建一个缓冲buffer
	p := make([]byte,num)
	n,err := reader.Read(p)	//读到p这个切片内，然后返回读了多少数据，以及可能存在的错误
	if n > 0{
		return p[:n],nil	//只返回读取的数量
	}

	return p,err
}

func sampleReadStdin(){
	fmt.Println("Please input from stdin:")
	data,_ := ReadFrom(os.Stdin,11)	//从控制台读取输入
	fmt.Println(data)
}

func sampleReadFromString(){
	data,_ := ReadFrom(strings.NewReader("from string"),12)	//从字符串读取

	fmt.Println(data)
}

func sampleReadFile(){			//从文件读取
	file,_ := os.Open("fib.txt")
	defer file.Close()
	data,_ := ReadFrom(file,9)
	fmt.Println(string(data))
}

func main(){
	//sampleReadFromString()
	//sampleReadStdin()
	sampleReadFile()
}
