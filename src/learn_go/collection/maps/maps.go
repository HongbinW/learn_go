package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"school": "xd",
		"degree": "master",
		"grade" : "second",
		"feel" : "bad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)		// m2 == empty map
	fmt.Println(m2)

	var m3 map[string]int			// m3 == nil
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k,v := range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	//degreeLevel, ok := m["degree"]
	//fmt.Println(degreeLevel,ok)
	//degreeLvel, ok := m["degee"]	//如果没有该键，则取出的是空
	//fmt.Println(degreeLvel,ok)

	//所以一般采取如下写法，如果取到值则ok=true，否则为false
	if degreeLevel, ok := m["degree"]; ok {
		fmt.Println(degreeLevel)
	}else{
		fmt.Println("key dose not exist")
	}

	fmt.Println("Delete values")
	delete(m,"feel")
	fmt.Println(m)

	fmt.Println(maxLengthSubStringWithoutDuplicate("abcabcbb"))
	fmt.Println(maxLengthSubStringWithoutDuplicate("bbbbb"))
	fmt.Println(maxLengthSubStringWithoutDuplicate("pwwkew"))
	fmt.Println(maxLengthSubStringWithoutDuplicate("p"))
	fmt.Println(maxLengthSubStringWithoutDuplicate(""))
	fmt.Println(maxLengthSubStringWithoutDuplicate("abcdef"))
	fmt.Println(maxLengthSubStringWithoutDuplicate("西安电子科技大学哈哈哈"))//中文是错的，是因为我们将字符串转为了byte，所以编码出错
	fmt.Println(maxLengthSubStringWithoutDuplicate("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))

}

//寻找最长不含有重复字符的子串

func maxLengthSubStringWithoutDuplicate(s string) int{
	//var maxLength int
	//var start int	//记录当前不重复的最开始的位置
	//var lastOccurred = make(map[byte]int)
	//for i, ch := range []byte(s){
	//	if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
	//		start = lastI + 1
	//	}
	//	if i - start + 1 > maxLength{
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i
	//}
	//return maxLength
	var maxLength int
	var start int	//记录当前不重复的最开始的位置
	var lastOccurred = make(map[rune]int)
	for i, ch := range []rune(s){
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
