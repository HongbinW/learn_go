package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int, op string) (int,error){
	switch op {
	case "+":
		return a+b, nil
	case "-":
		return a-b, nil
	case "*":
		return a*b, nil
	case "/":
		q, _  := div(a,b)
		return q, nil
		//return a/b
	default:
		//panic("upsupported operation :" + op)
		return 0, fmt.Errorf("unsupported operation: %s" , op)
	}
}

// 13 / 3 == 4 ... 1
func div(a,b int) (int, int){
		return a / b, a % b		////这种方式更为建议
		//q = a / b
		//r = a % b
		//return q,r
}

func apply(op func(int,int) int, a,b int) int{	//后面这俩int直接放到func里面了
	p := reflect.ValueOf(op).Pointer()		//通过反射的方式，获取函数名
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args " +
		"(%d, %d)" , opName, a, b)
	return op(a,b)
}

func pow(a,b int) int{
	return int(math.Pow(float64(a),float64(b)))
}

func sum(numbers ...int)int{
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func main() {
	//fmt.Println(eval(3,4,"/"))
	if result,err := eval(3,4,"%"); err != nil{
		fmt.Println("Error:" , err)
	}else{
		fmt.Println(result)
	}
	//fmt.Println(div(13,3))
	m,n := div(13,3)
	fmt.Println(m,n)

	fmt.Println(apply(pow,3,4))

	fmt.Println(apply(			//main.main.func1 定义匿名函数
		func(i int, i2 int) int {
		return int(math.Pow(float64(i),float64(i2)))
	},3,4))

	fmt.Println(sum(1,2,3,4))
}
