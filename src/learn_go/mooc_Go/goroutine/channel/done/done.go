package main

import (
	"fmt"
)

type worker struct {
	in chan int
	done chan  bool		//这种有点生产者消费者的意思
}

func doWorker(id int,w worker){	//多加一个channel用于通知
	for n:= range w.in{
		fmt.Printf("Worker %d received %c \n", id, n)
		go func() {		//并行的发
			w.done <- true
		}()
	}
}
func createWorker(id int) worker{
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id,w)
	return w
}

func chanDemo(){
	var chs [10] worker
	for i := 0; i < 10 ; i++ {
		chs[i] = createWorker(i)
	}

	for i := 0; i < 10 ; i++ {
		chs[i].in <- 'a' + i
		//<-chs[i].done
	}

	for i := 0; i < 10 ; i++ {
		chs[i].in <- 'A' + i
		//<-chs[i].done		//这样就顺序打印了
	}
	//time.Sleep(time.Millisecond)	//等一会，让goroutine执行完（不确定能否执行完）
	for _,worker := range chs{	//这样有个问题就是，大写的打不出来，因为前面小写的发出去对方接受后发出done还没人收（等待收），大写的又发了（等待发），就会dead lock
		<- worker.done
		<- worker.done
	}
}

func main() {
	chanDemo()
}
