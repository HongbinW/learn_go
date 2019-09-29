package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	计算1~200各个数的阶乘，并把各个数的阶乘放到map中，并显示出来，要求用goroutine完成
 */

var(
	m = make(map[int]int, 10)
	lock sync.Mutex		//全局互斥锁
)

func test(i int) {
	res := 1
	for n := 1; n <= i ; n ++  {
		res *= n
	}
	lock.Lock()
	m[i] = res	//concurrent map read and map write
	lock.Unlock()
}

func main(){
	m[0] = 1
	for i := 1; i <= 20 ; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for i, v := range m{
		fmt.Printf("%d! = %d\n",i,v)
	}
	lock.Unlock()
}