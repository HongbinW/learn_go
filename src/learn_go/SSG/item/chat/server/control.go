package main
//总控
import (
	"fmt"
	"io"
	"learn_go/SSG/item/chat/common"
	"learn_go/SSG/item/chat/server/classifyProcess"
	"learn_go/SSG/item/chat/util"
	"net"
)
type Processor struct {
	Conn net.Conn
}

//根据客户端发送消息种类不同，决定调用哪个来处理
func (this *Processor) ServerProcessMsg(msg *common.Message)(err error){
	switch msg.Type {
	case common.LoginMsgType:
		//处理登录请求
		up := &classifyProcess.UserProcess{
			this.Conn}
		err = up.ServerProcessLogin(msg)
	case common.RegisterMsgType:
		//处理注册请求

	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func (this *Processor) primaryProcess()(err error){

	//循环读客户端发送信息
	for {
		tf := &util.Transfer{
			Conn: this.Conn,
		}
		msg,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务端也要退出了")
				return err
			}
			fmt.Println("readPkg err",err)
			return err
		}
		fmt.Println("msg=",msg)
		err = this.ServerProcessMsg(&msg)
		if err != nil {
			fmt.Println("server classifyProcess message fail",err)
			return err
		}
	}
}


