package fib

// 1 1 2 3 5 8 13 21 34
func Fibonacci() func() int{
	pre := 0
	cur := 1
	return func() int{
		pre,cur = cur , pre + cur
		return pre
	}
}