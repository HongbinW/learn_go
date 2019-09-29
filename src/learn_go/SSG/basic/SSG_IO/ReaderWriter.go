package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file,err := os.OpenFile("fib.txt",os.O_RDWR|os.O_APPEND,0666)	//读写、追加
	defer file.Close()
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{	//读到文件末尾了
			break
		}
		fmt.Print(str)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString("Hello,BeiJing")
	writer.Flush()
}

//文件内容复制，俩文件已存在
func copy(){
	if message,err := ioutil.ReadFile("fib.txt");err != nil{
		fmt.Println(err)
		return
	}else{
		ioutil.WriteFile("test.txt",message,os.ModeAppend)
	}
}