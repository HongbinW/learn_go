package main

import "fmt"

func updateSlice(s []int){
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] =" ,arr[2:6])
	fmt.Println("arr[:6] =" ,arr[:6])



	s1 := arr[2:]
	fmt.Println("s1 =" ,s1)
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)	//Slice本身没有数据，是对底层array的一个view
	fmt.Println(s1)
	fmt.Println(arr)

	s2 := arr[:]
	fmt.Println("s2 =" ,s2)
	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending Slice")
	arr[0], arr[2] = 0 ,2
	s1 = arr[2:6]
	s2 = s1[3:5]	//s1[3],s1[4]，但是实际上直接去s1[4]是取不出来的
	fmt.Println("arr =",arr)
	fmt.Println("s1 =",s1)
	fmt.Println("s2 =",s2)
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d",s1,len(s1),cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d",s2,len(s2),cap(s2))

	s3 := append(s2,10)
	s4 := append(s3,11)		//s4和s5 view 是一个新的array,看的不再是arr了
	s5 := append(s4,12)
	fmt.Println("s3 =",s3)
	fmt.Println("s4 =",s4)
	fmt.Println("s5 =",s5)
	fmt.Printf("arr = %v, len(arr) = %d, cap(arr) = %d",arr,len(arr),cap(arr))







}
