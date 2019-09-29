package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1 1 2 3 5 8 13 21 34
func fibonacci() intGen{
	pre := 0
	cur := 1
	return func() int{
		pre,cur = cur , pre + cur
		return pre
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)	//string实现了read，因此直接用它的方法

	// TODO: incorrect if p is too small，做法可以是先把数据缓存起来
	return strings.NewReader(s).Read(p)	//写入p
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){		//没有起始条件和递增条件，类似while，go没有while
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	//fmt.Println(f())	//
	printFileContents(f)
}
