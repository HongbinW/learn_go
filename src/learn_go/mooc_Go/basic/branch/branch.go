package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if_else_test()
	fmt.Println(switch_test(3,5,"+"))
	fmt.Println(
		grade(85),
		grade(91),
		grade(63),
		grade(71),
		grade(101),	//panic会中断报错，其他分数也不会打印出来
		grade(0))
}

func if_else_test(){
	const filename = "abc.txt"
	//写法一
	contents, err := ioutil.ReadFile(filename)	//该函数有两个返回值，第一个是个byte数组，第二个是出错信息
	if err != nil {
		fmt.Print(err)
	}else{
		fmt.Printf("%s\n",contents)
	}

	//写法二
	if contents,err := ioutil.ReadFile(filename); err != nil{	//contents和err都是在if里定义的，因此他们的作用域就在if语句里
		fmt.Print(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
}

func switch_test(a,b int, op string) int{
	var result int			//switch会自动break，除非使用fallthrough
	switch op{		//switch后有变量
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsuported operator: " + op)
	}
	return result
}

func grade(score int) string{	//函数名grade，返回类型string
	g := ""
	switch {	//switch没有表达式，在case里加判断即可
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}


