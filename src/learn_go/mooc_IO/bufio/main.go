package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	strReader := strings.NewReader("hello world")

	bufReader := bufio.NewReader(strReader)

	data,_ := bufReader.Peek(5)	//返回缓存的一个切片，该切片引用缓存前n字节数据
	fmt.Println(string(data),bufReader.Buffered())	//同时返回bufReader中缓存了多少字符

	str,_ := bufReader.ReadString(' ')//用空格来切分，读取到空格为止，包含空格
	fmt.Println(str,bufReader.Buffered())	//返回剩下的字符

	// ReadString读取，Peek只读不取


	w := bufio.NewWriter(os.Stdout)	//默认输出到屏幕
	fmt.Fprint(w,"Hello ")	//Fprint写入到文件
	fmt.Fprint(w,"world!")	//写到bufio中
	w.Flush()	//提交到stdout中



}
