package main

import (
	"fmt"
	mock2 "learn_go/mooc_Go/interface/mock"
	real2 "learn_go/mooc_Go/interface/real"
	"time"
)

type Retriever interface {	//使用者定义接口，实现只需要去实现其具体的方法即可
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get(url)	//http要写
}

type Poster interface{
	Post(url string, form map[string]string) string
}

func post(poster Poster){
	poster.Post(url,
		map[string]string{
			"name":"whb",
			"course":"goland",
		})
}

type RetrieverPoster interface {	// 组合接口
	Retriever
	Poster
}

const url = "http://www.imooc.com"
func session(s RetrieverPoster) string{
	s.Post(url,map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriver := mock2.Retriever{"this is a fake imooc com"}
	r = &retriver
	//fmt.Printf("%T %v\n",r,r)	//打印类型  里面的值   mock.Retriever {this is a fake imooc com}
	inspect(r)
	r = &real2.Retriever{ // 修改为指针后  *real.Retriever &{Mozilla/5.0 1m0s}
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
	if realRetriver, ok := r.(*real2.Retriever);ok{
		fmt.Println(realRetriver.TimeOut)
	}else{
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriver))	//这个retriever是mockRetriever，同时实现了retriever和post方法，相当于有俩接口
}

func inspect(r Retriever) {
	fmt.Println("Inspect",r)
	fmt.Printf("%T %v\n",r,r)
	//Type switch
	switch v := r.(type) {
	case *mock2.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
