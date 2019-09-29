package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	//循环接收客户端发送的数据
	defer conn.Close()
	for{
		buf := make([]byte,1024)
		//读取客户端发送的消息，如果没有消息，则阻塞在这
		fmt.Printf("Server is waiting for client %s sending\n",conn.RemoteAddr().String())
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("server read err=",err)
			return
		}
		//if err == io.EOF {
		//	fmt.Println("客户端已退出")
		//	return
		//}
		//fmt.Printf("server read %d byte data\n",n)
		//显示到服务器终端
		fmt.Print(string(buf[:n]))	//不用打ln了，因为在客户端读入的时候，已经读\n了
		//一定要带长度n，因为后面不知道是啥，要显示真正读到的东西
	}
}

func main() {
	fmt.Println("Server start Listening...")
	//tcp：表示使用的网络协议
	//0.0.0.0:8888：表示在本地监听8888端口
	listen,err := net.Listen("tcp","127.0.0.1:8888")	//0.0.0.0 支持Ipv4和Ipv6,127.0.0.1只支持Ipv4
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()

	//循环等待客户端来连接
	for {
		fmt.Println("waiting for connection...")
		//等待连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err = ",err)
		}else{
			fmt.Println(conn.RemoteAddr())
		}
		go process(conn)
	}


}

//telnet测试
//在CMD中输入 telnet ip port：查看该ip的port是否在监听，如果是则进入，通过ctrl + ]退出