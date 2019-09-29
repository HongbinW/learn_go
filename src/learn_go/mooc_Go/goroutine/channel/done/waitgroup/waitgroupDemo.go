package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup	//使用同一个wg
}

func doWorker(id int,c chan int,wg *sync.WaitGroup){	//多加一个channel用于通知
	for n:= range c{
		fmt.Printf("Worker %d received %c \n", id, n)
		wg.Done()
	}
}
func createWorker(id int,wg *sync.WaitGroup) worker{
	w := worker{
		in:   make(chan int),
		wg:   wg,
	}
	go doWorker(id,w.in,wg)
	return w
}

func chanDemo(){

	var wg sync.WaitGroup
	//wg.Add(20)	//任务数
	var chs [10] worker
	for i := 0; i < 10 ; i++ {
		chs[i] = createWorker(i,&wg)
	}

	for i := 0; i < 10 ; i++ {
		wg.Add(1)	//可以随时加任务
		chs[i].in <- 'a' + i
	}

	for i := 0; i < 10 ; i++ {
		wg.Add(1)
		chs[i].in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
