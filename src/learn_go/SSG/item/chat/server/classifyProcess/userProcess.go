package classifyProcess

import (
	"encoding/json"
	"fmt"
	"learn_go/SSG/item/chat/common"
	"learn_go/SSG/item/chat/server/model"
	"learn_go/SSG/item/chat/util"
	"net"
)

//此处结构体的设计主要考虑方法用到哪些东西
type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(msg *common.Message)(err error) {
	var loginMsg common.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	//回复
	var resMsg common.Message
	resMsg.Type = common.LoginResMsgType

	var loginResMsg common.LoginResMsg

	user,err := model.MyUserDao.Login(loginMsg.UserId,loginMsg.PassWord)

	data, err := json.Marshal(loginResMsg)
	if err != nil{
		fmt.Println("json.Marshal fail",err)
		return
	}
	//封装完回复的message
	resMsg.Data = string(data)

	//对resMsg序列化，准备发送
	data,err = json.Marshal(resMsg)
	if err != nil{
		fmt.Println("json.Marshal fail",err)
		return
	}
	//为防止丢包，依然也要先发长度，再发数据
	transfer := &util.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	return
}



