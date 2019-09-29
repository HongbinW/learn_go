package classifyProcess

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"learn_go/SSG/item/chat/common"
	"learn_go/SSG/item/chat/util"
	"net"
)

type UserProcess struct {

}

//登录
func (this *UserProcess) Login(loginMsg common.LoginMsg) (err error){
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err:",err)
		return
	}
	defer conn.Close()
	//组装message
	var message common.Message
	message.Type = common.LoginMsgType

	data, err := json.Marshal(loginMsg)
	if err != nil{
		fmt.Println("json.Marshal err:",err)
		return
	}

	message.Data = string(data)

	//将message进行序列化，切片发送
	data, err = json.Marshal(message)
	if err != nil{
		fmt.Println("json.Marshal err:",err)
		return
	}
	//先将data长度发送给服务器，这样服务器收到时可以验证，防止丢包
	//获取data长度，将其转换为[]byte
	var pkg uint32
	pkg = uint32(len(data))
	dataLength := make([]byte,4)	//因为pkg是32位的，所以给他留四位
	binary.BigEndian.PutUint32(dataLength,pkg)
	n,err := conn.Write(dataLength)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(dataLength) err:",err)
		return
	}
	fmt.Println("发送数据的字节数正确",pkg)
	//发送要发送的数据
	_,err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write(data) err:",err)
		return
	}

	//还需要接收服务器端返回的消息
	tf := &util.Transfer{Conn:conn}
	msg,err := tf.ReadPkg()
	if err != nil {
		fmt.Println("read pkg fail",err)
		return
	}

	var loginResMsg common.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data),&loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal received data fail",err)
		return
	}
	if loginResMsg.Code == common.Success {
		fmt.Println("登录成功...")

		//需要起一个协程，保持与服务端的通讯，接收服务器推送消息
		go processServerMsg(conn)
		//显示登陆后的菜单
		for{
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMsg.Error)

	}

	return
}

//与服务端保持通讯
func processServerMsg(conn net.Conn) {
	tf := &util.Transfer{
		Conn: conn,
	}
	for{
		fmt.Println("客户端正在读取服务器发送的消息")
		msg,err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.readPkg err:",err)
			return
		}
		fmt.Println(msg)
	}
}