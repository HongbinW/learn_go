package main

import (
	"bufio"
	"fmt"
	fib2 "learn_go/mooc_Go/functional/fib"
	"os"
)

func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)	//defer相当于一个栈，先进后出
	fmt.Println(3)			//defer不怕他中间return,panic

	for i := 0; i < 100; i ++{
		defer fmt.Println(i)
		if(i == 30){
			panic("print too many")
		}
	}
}

func writeFile(filename string){
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)	//先写到内存里，这样比直接写更快
	defer writer.Flush()

	f := fib2.Fibonacci()
	for i := 0; i < 20 ; i++  {
		fmt.Fprintln(writer,f())
	}
}


func main(){
	tryDefer()
	//writeFile("fib.txt")
}

