package main

import "github.com/garyburd/redigo/redis"

var Pool redis.Pool

func init(){
	Pool = redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp","0.0.0.0:6379")
		},
		MaxIdle:         8,		//最大空闲连接数
		MaxActive:       16,		//最大连接数
		IdleTimeout:     100,		//最大空闲时间
	}
}

func main() {
	p := Pool.Get()		//获取连接

	Pool.Close()		//关闭连接池
}
