package main

import (
	"fmt"
	"learn_go/SSG/item/chat/client/classifyProcess"
	"learn_go/SSG/item/chat/common"
	"os"
)

//定义协议

func main() {
	//conn,err := net.Dial("tcp","0.0.0.0:8889")
	//if err != nil {
	//	fmt.Println("tcp connect err:",err)
	//	return
	//}

	var loginMsg common.LoginMsg
	var key int
	//var loop bool = true
	for {
		fmt.Println("------------欢迎登陆多人聊天系统------------")
		fmt.Println("               1.登陆聊天系统")
		fmt.Println("               2.注册用户")
		fmt.Println("               3.退出系统")
		fmt.Println("请选择（1-3）")
		fmt.Println("--------------------")
		fmt.Scanln(&key)
		switch key {
			case 1:
				fmt.Println("登录聊天室")
				fmt.Println("请输入用户Id")
				fmt.Scanln(&loginMsg.UserId)
				fmt.Println("请输入密码")
				fmt.Scanln(&loginMsg.PassWord)

				userProcess := &classifyProcess.UserProcess{}
				userProcess.Login(loginMsg)
				//loop = false
			case 2:
				fmt.Println("用户注册")
				//loop = false
			case 3:
				fmt.Println("退出系统")
				//loop = false
				os.Exit(0)
			default:
				fmt.Println("输入有误,请重新输入")
		}
	}

}
