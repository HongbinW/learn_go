package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn){
	for{
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("received successfully")
			return
		}
		str := string(buf[:n])
		fmt.Print(str)
	}
}

func main() {
	for{
		listen, err := net.Listen("tcp", "0.0.0.0:8888")
		if err != nil{
			fmt.Println("listen err:",err)
			continue
		}
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:",err)
			continue
		}
		go process(conn)
	}
}
