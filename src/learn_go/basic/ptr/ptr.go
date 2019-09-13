package main

import "fmt"

func swap1(a,b int) (int,int){
	return b,a
}
func swap2(a,b *int){
	*b,*a = *a,*b
}

func main() {
	a, b := 3,4
	a,b = swap1(a,b)
	swap2(&a,&b)
	fmt.Println(a,b)
}
