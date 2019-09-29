package familiyAccount

import "fmt"

type FamilyAccount struct {
	key string

	//控制是否退出for
	loop bool
	//账户余额
	balance float64
	//每次收支金额
	money float64
	//收支说明
	note string
	//收支详情，收支直接对该变量进行拼接
	detail string
	//是否有收支行为
	flag bool
}

//构造器
func NewFamiliyAccount() *FamilyAccount{
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		detail:  "收支\t账户金额\t收支金额\t说明\n",
		flag:    false,
	}
}


func (f *FamilyAccount) Mainmenu(){
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1 收支明细")
		fmt.Println("                 2 登记收入")
		fmt.Println("                 3 登记支出")
		fmt.Println("                 4 退出")
		fmt.Println("                 请选择（1-4）：")

		fmt.Scanln(&f.key)
		switch f.key {
		case "1":
			f.showDetails()
		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
		default:
			fmt.Println("请输入正选的选项..")
		}
		if !f.loop {
			break
		}
	}
}

func (f *FamilyAccount) showDetails(){
	fmt.Println("-----------------当前收支明细记录-----------------")
	if !f.flag{
		fmt.Println("来一笔？")
	}else {
		println(f.detail)
	}
}

func (f *FamilyAccount) income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&f.money)
	fmt.Println("本次收入说明：")
	fmt.Scanln(&f.note)
	f.balance += f.money
	f.detail += fmt.Sprintf("收入\t%v\t%v\t%v\n",f.balance,f.money,f.note)
	f.flag = true
}

func (f *FamilyAccount) pay(){
	fmt.Println("-----------------登记支出-----------------")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足")
		return
	}
	fmt.Println("本次支出说明：")
	fmt.Scanln(&f.note)
	f.balance -= f.money
	f.detail += fmt.Sprintf("支出\t%v\t%v\t%v\n",f.balance,f.money,f.note)
	f.flag = true
}

func (f *FamilyAccount) exit(){
	fmt.Println("确定要退出吗？y/n")
	valid := ""
	for{
		fmt.Scanln(&valid)
		if valid == "y" || valid == "n"{
			break
		}
		fmt.Println("输入有误，重新输入 y/n")
	}
	if valid == "y" {
		f.loop = false
	}
}
