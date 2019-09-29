package main

import (
	"fmt"
	"math"
	"testing"
)

func calcTriangle(a,b int) int{
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}


func triangle(){
	var a,b int = 3,4
	fmt.Println(calcTriangle(a,b))
}

func TestTriangle(t *testing.T){	//可以直接运行
	tests := []struct{a,b,c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
	}

	for _,tt := range tests{
		if actural := calcTriangle(tt.a, tt.b); actural != tt.c {
			t.Errorf("calcTriangle(%d,%d);" + "got %d; expected %d",tt.a,tt.b,actural,tt.c)
		}
	}
}
