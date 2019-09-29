package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"learn_go/SSG/item/chat/server/model"
	"net"
	"time"
)

var pool *redis.Pool

func initPool(){
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp","0.0.0.0:6379")
		},
		MaxIdle:         0,
		MaxActive:       16,
		IdleTimeout:     300 * time.Second,
	}
}

func process(conn net.Conn){
	defer conn.Close()
	//创建总控
	ps := &Processor{
		Conn:conn,
	}
	err := ps.primaryProcess()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("process fail",err)
		return
	}
}

func initUserDao(){
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool()
	initUserDao()
	fmt.Println("Server is listening at 8889")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil {
		fmt.Println("Listen err:",err)
		return
	}
	defer listen.Close()
	for{
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Conn err:",err)
			return
		}else {
			//一旦连接成功，则启动一个协程与客户端保持通讯
			go process(conn)
		}
	}
}
