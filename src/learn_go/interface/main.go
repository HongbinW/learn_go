package main

import (
	"fmt"
	"learn_go/interface/mock"
	"learn_go/interface/real"
	"time"
)

type Retriever interface {	//使用者定义接口，实现只需要去实现其具体的方法即可
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc com"}
	//fmt.Printf("%T %v\n",r,r)	//打印类型  里面的值   mock.Retriever {this is a fake imooc com}
	inspect(r)
	r = &real.Retriever{			// 修改为指针后  *real.Retriever &{Mozilla/5.0 1m0s}
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	//fmt.Printf("%T %v\n",r,r)	//real.Retriever { 0s} 第一个是空格，第二个是0秒。刚好对应struct属性
	inspect(r)


	//三种方式得到interface内含的值：
	//fmt.Printf("%T %v\n",r,r)
	//Type switch
	//Type assertion


	//Type assertion
	if realRetriver, ok := r.(*real.Retriever);ok{
		fmt.Println(realRetriver.TimeOut)
	}else{
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n",r,r)
	//Type switch
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
