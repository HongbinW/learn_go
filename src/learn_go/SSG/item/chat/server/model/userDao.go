package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var(
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{pool:pool}
	return
}

func (this *UserDao) getUserById(conn redis.Conn,id string)(user *User, err error){
	res, err := redis.String(conn.Do("HGet","users",id))
	if err != nil {
		if err == redis.ErrNil {	//没查到
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	err = json.Unmarshal([]byte(res),&user)
	if err != nil{
		fmt.Println("json.Unmarshal fail",err)
		return
	}
	return
}

//登录校验
//验证正确返回user实例
//任意有误，返回error
func (this *UserDao) Login(userId,userPwd string)(user *User, err error){
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil {
		return
	}
	if userPwd != user.PassWord {
		err = ERROR_USER_WRONGPWD
		return
	}
	return
}