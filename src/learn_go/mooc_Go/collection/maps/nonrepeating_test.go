package main

import "testing"

func TestSubstr(t *testing.T){
	tests := []struct{
		s string
		ans int
	}{
		{"abcabcbb",3},
		{"bbbbb",1},
		{"pwwkew",3},
		{"p",1},
		{"",0},
		{"abcdef",6},
		{"西安电子科技大学哈哈哈",9},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花",8},
	}

	for _,tt := range tests{
		if actural := maxLengthSubStringWithoutDuplicate(tt.s);actural != tt.ans{
			t.Errorf("maxLengthNoRepeating string is %s, got %d, expected %d",tt.s,actural,tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B){
	s,ans :="黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花",8
	for i := 0; i < b.N ; i++{	//具体算多少遍，系统会给出结果
		if actural := maxLengthSubStringWithoutDuplicate(s); actural != ans {
			b.Errorf("maxLengthNoRepeating string is %s, got %d, expected %d", s, actural, ans)
		}
	}
}