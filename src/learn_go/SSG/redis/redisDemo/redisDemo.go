package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn,err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("dial err:",err)
		return
	}
	fmt.Println("connect success",conn)
	defer conn.Close()

	_, err = conn.Do("Set","name","tom")	//命令
	if err !=nil {
		fmt.Println("set err:",err)
		return
	}
	_,err = conn.Do("Set","age","18")
	if err !=nil {
		fmt.Println("set err:",err)
		return
	}
	s ,err := redis.String(conn.Do("Get", "name"))	//加s返回多个字符串
	t ,err := redis.String(conn.Do("Get", "age"))

	o ,err := redis.String(conn.Do("Get", "whn"))
	fmt.Println("not exit",o,err)
	if err != nil {
		fmt.Println("no value",o)
	}else{
		fmt.Println("has value",o)
	}
	//返回类型是interface{}
	//因为name对应的值是string
	//如上使用redis.string	redis.strings	redis.Int

	fmt.Println(s)
	fmt.Println(t)
}
