package classifyProcess

import (
	"fmt"
	"os"
)

//保持跟服务端的通讯


//显示登陆成功后的界面
func ShowMenu(){

	fmt.Println("欢迎,***用户")
	fmt.Println("1.显示在线用户列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.信息列表")
	fmt.Println("4.退出系统")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误，情重新输入")
	}

}

