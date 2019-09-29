package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//有问题，读取的序列化字符串无法反序列化
//已解决：反序列化时，一定要确定切片实际长度，不能比实际长

type Monster struct{
	Name string
	Age int
	Skill string
}

func (m *Monster) Store(){
	data,err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化错误,err：",err)
		return
	}
	file,err := os.OpenFile("test.txt",os.O_APPEND|os.O_WRONLY,0666)
	if err != nil {
		fmt.Println("打开文件错误,err：",err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
}

func (*Monster) ReStore() *Monster{
	file,err := os.OpenFile("test.txt",os.O_RDONLY,0666)
	if err != nil {
		fmt.Println("打开文件错误,err：",err)
	}
	reader := bufio.NewReader(file)
	str,err := reader.ReadBytes('\n')
	fmt.Println(string(str))
	//{"Name":"JZOFFER","Age":"2","Skill":"Exercise"}
	//{"Name":"JZOffer","Age":2,"Skill":"Exercise"}
	var mon Monster
	json.Unmarshal(str,&mon)
	return &mon
}

func main(){
	monster := Monster{
		Name:  "JZOffer",
		Age:   2,
		Skill: "Exercise",
	}
	monster.Store()
	var m *Monster
	m = monster.ReStore()
	fmt.Printf("Monster,%v",&m)


	//str := "{\"Name\":\"JZOFFER\",\"Age\":\"2\",\"Skill\":\"Exercise\"}"
	//fmt.Println(str)
	//var b []byte = []byte(str)
	//var monster Monster
	//json.Unmarshal(b,&monster)
	//fmt.Printf("反序列化后 monster=%v monster.Name=%v \n",monster,monster.Name)


	//byte,_ := json.Marshal(monster)
	//var monkey Monster
	//json.Unmarshal(byte,&monkey)
	//fmt.Println(monkey)

}