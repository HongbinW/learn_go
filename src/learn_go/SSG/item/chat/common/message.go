package common

//定义协议

const(
	LoginMsgType = "LoginMsg"
	LoginResMsgType = "LoginResMsg"
	RegisterMsgType = "RegisterMsg"
	UserNotExist = 500
	Success = 200
	WrongPwd = 400
)

type Message struct {
	Type string `json:"type"`		//消息类型	//指定序列化后的名字
	Data string `json:"data"`		//消息内容
}

type LoginMsg struct {
	UserId string `json:"user_id"`
	PassWord string `json:"pass_word"`
	UserName string	`json:"user_name"`
}

type LoginResMsg struct {
	Code int `json:"code"`		//状态码： 503 表示用户未注册	200表示登录成功
	Error string `json:"error"`	//返回错误信息
}

type RegisterMsg struct {

}