package model

type User struct {
	//一定要让序列化后的字符串与结构体字段的tag名字一样！！！
	UserId string `json:"user_id"`
	PassWord string `json:"pass_word"`
	UserName string	`json:"user_name"`
}