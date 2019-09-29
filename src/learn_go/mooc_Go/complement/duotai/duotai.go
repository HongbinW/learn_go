package main

import "fmt"

//多态

type Sleep struct {
	start, end int
}

type Wake struct {
	morning,evening int

}

type Schedule interface {
	HowLong() int
}


func (s Sleep) Diff() int{
	return s.end - s.start
}

func (s Sleep) HowLong() int{
	return s.end - s.start
}

func (w Wake) HowLong() int{
	return w.evening - w.morning
}

func process(s Schedule) int{
	return s.HowLong()
}

func main(){
	s := &Sleep{1,2}
	w := Wake{3,10}
	fmt.Println(process(s))
	fmt.Println(process(w))
	fmt.Printf("%T %v \n",s,s)
	fmt.Printf("%T %v \n",w,w)

	var ss Schedule		//接口变量
	ss = s
	ss = w
	fmt.Printf("%T %v \n",ss,ss)


}