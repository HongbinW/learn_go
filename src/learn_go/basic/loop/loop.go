package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	file,err := os.Open(filename)		//op.Open()读文件
	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)

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

	printFile("abc.txt")
	forever()
}
