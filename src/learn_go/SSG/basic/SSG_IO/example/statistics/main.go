package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//每读取一行，就统计该行有多少个英文、数字、空格和其它字符

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}


func main() {
	file,_ := os.Open("C:\\Users\\admin\\Desktop\\租房信息整理.txt")
	//file,_ := os.Open("fib.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	index := 1
	for{
		str,err := reader.ReadString('\n')		//TODO:有问题，结果是byte[]
		if err == io.EOF {
			break
		}
		var count CharCount
		for _,v := range str{
			switch {
			case v >= '0' && v <= '9':
				count.NumCount++
			case v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
		fmt.Printf("第%d行中的字母有%d个，数字有%d个，空格有%d个，其它字符有%d个\n",index,count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
		index ++
	}
}
