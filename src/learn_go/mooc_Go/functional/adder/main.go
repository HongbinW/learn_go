package main

import "fmt"

func adder() func(int) int{	//返回值是一个函数
	sum := 0
	return func(v int) int{
		sum += v
		return sum
	}
}

//以下是一个较为正统的函数式编程，只有函数和常数
type iAdder func(int) (int,iAdder)

func adder2(base int) iAdder{
	return func(v int)(int,iAdder){
		return base + v, adder2(base + v)
	}
}




func main() {
	//a := adder()	//a中存了sum
	//for i := 0; i < 10 ; i++ {
	//	fmt.Printf("0 + 1 + ... + %d = %d\n ", i ,a(i))
	//}

	a := adder2(0)
	for i := 0; i < 10 ; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n ", i ,s)
	}
}
