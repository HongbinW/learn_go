package main

import (
	"errors"
	"fmt"
)

func tryRecover(){
	defer func(){
		r := recover()
		if err,ok := r.(error); ok{		//转成一个error
			fmt.Println("Error occured:",err)
		}else{
			panic(r)
		}
	}()
	//panic(errors.New("this is my error"))
	//
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	defer func() {
		panic(errors.New("defer panic"))
	}()
	panic(errors.New("test panic"))
}

func main() {
	tryRecover()
}
