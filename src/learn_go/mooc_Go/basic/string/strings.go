package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello!哈喽"		//中文三字节,UTF-8
	fmt.Println(len(s))
	for _,b := range []byte(s){
		fmt.Printf("%X " ,b)
	}
	fmt.Println()

	// string 先做utf-8解码，转完后，又转unicode，又放到rune四字节里
	for i,ch := range s{		//ch is rune
		fmt.Printf("(%d %c)",i,ch)
		//fmt.Printf("(%d %X)",i,ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0{
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i,ch := range []rune(s){
		fmt.Printf("(%d %c)",i,ch)
	}


	str := strings.Fields("  a b  d e   ")
	fmt.Println(str)
	t := "abc"
	fmt.Println(strings.Index(t,"bc"))




}
