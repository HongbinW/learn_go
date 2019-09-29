package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	line,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString from stdin,err=",err)
	}
	//发送给line
	n,err := conn.Write([]byte(line))
	if err != nil{
		fmt.Println("conn.write err=",err)
	}
	fmt.Printf("client has send %d byte data\n",n)

}
