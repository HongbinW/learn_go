package main

import "fmt"

// 统计1~200000，哪个是素数

func isValid(n int) bool{
	for i := 2; i <= n/2 ; i++{
		if n % i == 0 {
			return false;
		}
	}
	return true;
}

func process(numChannel chan int, c chan int, b chan bool){
	for {
		n,ok := <-numChannel
		if !ok {	//如果取不到值了，直接退出
			break
		}
		if isValid(n) {
			c <- n
		}
	}
	b <- true
}

func putNum(n int,c chan int){
	for i := 1; i <= n ; i++ {
		c <- i
	}
	close(c)
}


func main(){
	putChannel := make(chan int,1000)	//将所有数字放到putChannel
	go putNum(200000,putChannel)
	c := make(chan int,2000)	//放入结果
	b := make(chan bool,4)

	//开启四个协程，判断是否是素数，如果是，则放入c
	for i := 0; i < 4 ; i ++ {
		go process(putChannel,c,b)
	}

	go func() {
		//看是否完成b中是否有四个T
		for i := 0; i < 4; i ++{
			<- b		//取不到四个就在此等待
		}
		close(c)
	}()

	for v := range c{
		fmt.Println(v)
	}

}