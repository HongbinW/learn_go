package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("%v ,len = %d, cap = %d \n",s,len(s),cap(s))
}

func main() {
	var s []int		//Zero value for slice is nil
	for i := 0 ; i < 100 ; i ++{
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int,16)	//len = 16

	s3 := make([]int,10,32)	//len = 10,cap = 32

	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2,s1)		//先目标切片，后源切片
	printSlice(s2)

	fmt.Println("Delete elements from slice")
	//删掉下标为3的元素
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)
	fmt.Println("front",front)
	fmt.Println("Poping from tail")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2)-1]
	printSlice(s2)
	fmt.Println("tail",tail)
}
