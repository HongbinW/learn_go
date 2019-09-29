package util

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"learn_go/SSG/item/chat/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

//读取消息，返回message
func (this *Transfer) ReadPkg()(msg common.Message, err error){
	//buf := make([]byte,8096)
	_,err = this.Conn.Read(this.Buf[:4])		//在conn没有被关闭的情况下才会阻塞，如果任意一方关闭，则不会阻塞
	if err != nil{
		return
	}
	//根据buf[:4]转成一个uint32类型，得到数据长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据 pkgLen读取消息内容
	n,err := this.Conn.Read(this.Buf[:pkgLen])		//从connect中读取pkgLen长的字节到buf中
	if n != int(pkgLen) || err != nil{
		return
	}

	//反序列化
	err = json.Unmarshal(this.Buf[:pkgLen],&msg)		//msg一定要带&！！！！！！	//反序列化时长度一定要与序列化长度一致
	if err != nil {
		err = errors.New("json.Unmarshal err")
		return
	}

	return
}

//接收要发送数据的[]byte
func (this *Transfer) WritePkg(data []byte)(err error)  {
	//先发长度，再发数据
	//buf := make([]byte,4)
	binary.BigEndian.PutUint32(this.Buf[:4],uint32(len(data)))
	n,err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("send resMsgLength fail",err)
		return
	}
	n,err = this.Conn.Write(data)
	if n != len(data) || err != nil {
		fmt.Println("send resMsg fail",err)
		return
	}
	return
}
