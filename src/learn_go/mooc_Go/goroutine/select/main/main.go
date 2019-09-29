package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)	//随机1500毫秒
			out <- i
			i ++
		}
	}()
	return out
}
func doWorker(id int,c chan int){	//多加一个channel用于通知
	for n:= range c{
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n", id, n)
	}
}
func createWorker(id int) chan<- int{
	in := make(chan int)
	go doWorker(id,in)
	return in
}
func main() {
	var c1,c2 = generator(),generator()
	w := createWorker(0)
	//var w chan int	// w is nil channel

	tm := time.After(10 * time.Second)	//运行10秒钟，会送一个channel
	tick := time.Tick(time.Second)	//定时1s，可以检查系统状态


	var values []int
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values,n)
		case n := <-c2:
			values = append(values,n)
		case activeWorker <- activeValue:	//如果是个nil，可以运行，但永远不会被select到
			values = values[1:]
		case <- time.After(800 * time.Millisecond):	//每次要select的时候，超过800毫秒没收到数据，执行
			fmt.Println("timeout")
		case <- tick:	//每隔一秒返回当前队列长度
			fmt.Println("queue len =",len(values))
		case <- tm:
			fmt.Println("Bye")
			return
		}
	}
}

//问题：如果收到的数据快，接受数据会漏数据
//解决：把n存起来