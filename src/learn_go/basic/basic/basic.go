package main //必须main包？并且main包不用跟包名一致？  对！

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3		//用于包内部，不是全局变量
var ss = "kkk"	//函数外面不可以用:=来定义变量。必须以关键字开始

var(			//省略形式
	cc = 4;
	dd = "mmm"
	ee = true
)

func variableZeroValue(){	//定义变量
	var a int		//变量  变量名  变量类型
	var s string
	fmt.Printf("%d %q \n",a,s)	//printf,%q可以打印空串，q-->quotation 引用
}

func variableInitialValue(){
	var a,b int = 3, 4	//不能有未被使用的变量
	var s string = "abc"
	fmt.Println(a,s,b)
}

func variableTypeDeduction(){
	var a, b, c, d = 1,2,"abc",true	//自动判断类型
	fmt.Println(a,b,c,d)
}

func variableShorter(){
	a,b,c,d := 3,4,false,"def" // := 定义变量
	b = 5		//之后再赋值就不可以用:=了，即不能重复定义变量
	fmt.Println(a,b,c,d)

}

func main()  {
	//fmt.Println("Hello World~")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(cc,dd,ee)
	//var x rune = 'a'
	//fmt.Println(x)
	euler()
	triangle()
	var mm uint8	//范围0~255
	fmt.Println(mm)
	consts()
	enums()
}

func euler(){
	//fmt.Println(cmplx.Pow(math.E,1i * math.Pi) + 1)	//1i表示虚数i
 	//fmt.Println(cmplx.Exp((1i * math.Pi)) + 1)
	fmt.Printf("%.3f\n",cmplx.Exp((1i * math.Pi)) + 1)
}

func triangle(){
	a,b := 3,4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

const filename2 = "abc2.txt"
func consts(){
	const filename = "abc.txt"
	const(
		t = "tencent"	//go常量一般不大写，大写有自己的含义
		h = "huawei"
	)
	const a,b = 3,4	//a,b没确定类型，既可作认为是int 又可以认为是float
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename,c)
}

func enums(){
	const(
		cpp = iota
		python
		golang
		javaScript
	)
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,javaScript,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

