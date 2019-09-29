package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//把整数转换为二进制字符串
func convertToBin(n int) string{
	result := ""
	for ; n > 0; n = n >> 1{
		lsb := (n & 1)
		result = strconv.Itoa(lsb) + result  //数字转字符串函数Itoa
	}
	return result
}

func printFile(filename string){
	file,err := os.Open(filename)		//os.Open()读文件
	if err != nil{
		panic(err)
	}

	printFileContents(file)
	//scanner := bufio.NewScanner(file)
	//
	//for scanner.Scan(){		//没有起始条件和递增条件，类似while，go没有while
	//	fmt.Println(scanner.Text())
	//}
}

//Reader、Writer补充
func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){		//没有起始条件和递增条件，类似while，go没有while
		fmt.Println(scanner.Text())
	}
}

func forever(){
	for{						//也没有终止条件，什么都不加就是死循环
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		)

	printFile("basic/branch/abc.txt")

	s := `abc " d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))		//使用Reader后，适用范围更广了，不仅可以打印文件，也可以直接打印字符串

	//forever()
}
