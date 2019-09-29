package main

import "fmt"


func main() {

	//接收用户输入选项
	key := ""

	//控制是否退出for
	loop := true

	//账户余额
	balance := 10000.0
	//每次收支金额
	money := 0.0
	//收支说明
	note := ""
	//收支详情，收支直接对该变量进行拼接
	detail := "收支\t账户金额\t收支金额\t说明\n"
	//是否有收支行为
	flag := false

	for{
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1 收支明细")
		fmt.Println("                 2 登记收入")
		fmt.Println("                 3 登记支出")
		fmt.Println("                 4 退出")
		fmt.Println("                 请选择（1-4）：")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("-----------------当前收支明细记录-----------------")
			if !flag{
				fmt.Println("来一笔？")
			}else {
				println(detail)
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			balance += money
			detail += fmt.Sprintf("收入\t%v\t%v\t%v\n",balance,money,note)
			flag = true
		case "3":
			fmt.Println("-----------------登记支出-----------------")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			balance -= money
			detail += fmt.Sprintf("支出\t%v\t%v\t%v\n",balance,money,note)
			flag = true
		case "4":
			fmt.Println("确定要退出吗？y/n")
			valid := ""
			for{
				fmt.Scanln(valid)
				if valid == "y" || valid == "n"{
					 break
				}
				fmt.Println("输入有误，重新输入 y/n")
			}
			if valid == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正选的选项..")
		}
		if !loop {
			break
		}
	}

	fmt.Println("已退出家庭收支记账软件..")
}
