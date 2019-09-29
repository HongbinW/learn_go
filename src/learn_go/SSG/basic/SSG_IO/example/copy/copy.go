package main

//拷贝一个图片/电影/mp3拷贝到另一个文件
//C:\Users\admin\Desktop\Lena.jpg --> C:\Users\admin\Desktop\Lena2.jpg
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	srcPath := "C:\\Users\\admin\\Desktop\\Lena.jpg"
	destPath := "C:\\Users\\admin\\Desktop\\Lena2.jpg"
	_,err := FileCopy(srcPath,destPath)
	if err == nil {
		fmt.Println("拷贝完成")
	}else{
		fmt.Printf("拷贝错误%v\n",err)
	}
}

func FileCopy(srcPath,destPath string) (written int64, err error){
	src,errS := os.OpenFile(srcPath,os.O_RDONLY,0666)
	dest,errD := os.OpenFile(destPath,os.O_CREATE|os.O_WRONLY,0666)
	defer src.Close()
	defer dest.Close()
	if errS != nil  {
		fmt.Println("srcPath is Wrong")
		return
	}
	if errD != nil {
		fmt.Println("destPath is Wrong")
	}
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dest)
	return io.Copy(writer,reader)
}

func PathExists(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	} else if os.IsNotExist(err){
		return false,nil
	} else{
		return false,err
	}
}