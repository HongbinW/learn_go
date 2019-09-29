package main

import (
	"fmt"
	"strconv"
	"time"
)

//在主线程中，开启一个goroutine，该协程每隔一秒输出“hello world"
//主线程中每个一秒输出"hello golang"，输出10次后退出程序
//要求主线程和goroutine同时执行

func test(){
	for i := 0; i < 10 ; i ++ {
		fmt.Println("test hello world" + strconv.Itoa(i))
		time.Sleep(time.Millisecond)
	}
}
func main() {
	go test()
	for i := 0; i < 10 ; i ++ {
		fmt.Println("main hello golang" + strconv.Itoa(i))
		time.Sleep(time.Millisecond)
	}
}
