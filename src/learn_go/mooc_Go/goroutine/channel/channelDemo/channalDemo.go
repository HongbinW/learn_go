package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int){
		//for {
		//	if n, ok := <-c;!ok{
		//		break
		//	}else {
		//		fmt.Printf("Worker %d received %c \n", id, n)
		//	}
		//}
		for n:= range c{		//也可以实现上述代码的效果
			fmt.Printf("Worker %d received %c \n", id, n)
		}
}


func createWorker(id int) chan<- int{	//告诉用户，返回的channel只能发数据		// <-chan则只能收数据
	c := make(chan int)
	//go func() {		//必须要在goroutine中对channel处理，
	//	for {
	//		fmt.Printf("Worker %d received %c \n",id,<-c)
	//	}
	//}()
	go worker(id,c)
	return c
}

func chanDemo(){
	//var c chan int	//一个传int的channel			c == nil
	var chs [10]chan<- int
	for i := 0; i < 10 ; i++ {
		//chs[i] = make(chan int)
		//go worker(i,chs[i])
		chs[i] = createWorker(i)
	}
	//c <- 1		//需要别的goroutine来从这个channel里读数据
	for i := 0; i < 10 ; i++ {
		chs[i] <- 'a' + i
	}
	for i := 0; i < 10 ; i++ {
		chs[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)	//等一会，让goroutine执行完（不确定能否执行完）
}

func bufferedChannel(){
	c:= make(chan int,3)	//添加channel缓冲区

	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'b'
	c <- 'd'	//此时会deadlock
	time.Sleep(time.Millisecond)
}

func channelClose(){
	c:= make(chan int)

	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'b'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}


func main() {

	chanDemo()
	//bufferedChannel()
	//channelClose()
}
