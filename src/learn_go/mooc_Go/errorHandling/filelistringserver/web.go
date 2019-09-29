package main

import (
	filelist2 "learn_go/mooc_Go/errorHandling/filelistringserver/filelist"
	"log"
	"net/http"
	"os"
)

//获取错误
type appHandler func(writer http.ResponseWriter, request *http.Request) error

//统一处理错误
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic
		defer func() { //防止出现意外的错误，导致网页无法显示
			if r := recover(); r != nil {
			log.Printf("Panic :%v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer,request)
		if err != nil {
			log.Println("Error handling request:",err)	//log.Warn()

			// user error
			if userErr, ok := err.(userError); ok{		//如果是userError，则在此处判断并返回
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer,http.StatusText(code),code)
		}
	}
}
//定义一组可以给用户看到的错误，要不然不是HTTP常规错误，就全都是500了
type userError interface {
	error
	Message() string
}

//url写入一个地址，返回该地址文件的内容
func main() {
	http.HandleFunc("/",	//根路径
		errWrapper(filelist2.HandleFileList)) //filelist.HandleFileList是appHandler类型

	err := http.ListenAndServe(":8888",nil)
	if	err != nil{
		panic(err)
	}
}
