package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

//读文件头

func main(){
	file,_ := os.Open("lena.bmp")
	defer file.Close()
	var headA,headB byte	//一个word对应俩byte
	binary.Read(file,binary.LittleEndian,&headA)//字节序：Windows/linux字节序一般都是小端binary.LittleEndian
	binary.Read(file,binary.LittleEndian,&headB)//字节序：Windows/linux字节序一般都是小端binary.LittleEndian

	var size uint32
	binary.Read(file,binary.LittleEndian,&size)

	fmt.Printf("%c%c",headA,headB)


	// 直接使用结构体，更为简洁
	infoHeader := new(BitmapInfoHeader)
	binary.Read(file,binary.LittleEndian,infoHeader)	//通过反射填入结构体

	fmt.Println(infoHeader)
}

type BitmapInfoHeader struct {
	Size uint32
	Width int32
	Height int32
	//.....还有很多
}