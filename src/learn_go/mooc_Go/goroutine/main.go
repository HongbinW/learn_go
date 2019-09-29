package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i ++{	//不限goroutine数量
		go func(i int) {		//直接写成匿名函数，也可以写成调用函数
			for {
				//fmt.Printf("Hello from goroutine %d \n",i)		//直接使用i不安全，可以使用参数传递
				a[i]++		//print是IO操作，中间的等待时间会进行协程切换，而a[i]是对内存操作，不会进行协程切换
				runtime.Gosched()	//手动交出控制权
				// 此处主动释放不会永远停在这，因为a[i]和sleep都没有释放控制权，那如果此处执行完，执行sleep是没问题的
				// 但是如果，此处没执行完，是切换到main执行sleep，那再切回来的时候，会永远停在这，因为切不出去了
				// 其实就是看内部goroutine和main谁先拿控制权
			}
		}(i)		//使用参数传递不直接使用i：因为如果是用外部的i，内部的goroutine会跟main之间做切换，当i加到10时，会跳出来，然而如果之前是从内部切换到main的
					//那此时main已经加到10了，此时再切换回内部goroutine，会报超出数组范围异常
	}
	time.Sleep(time.Microsecond)		//未打印原因：还没这些goroutine还没来得及打印，main已经执行完i的递增，所以提前结束
										// main函数也是个goroutine，由于不交出控制前 ，会一直停在此处
	fmt.Println(a)	//由于a[i]操作不会进行协程切换，因此进入一个协程后会一直在里面，出不来
}