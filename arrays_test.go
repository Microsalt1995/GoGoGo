package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// 数组默认值是零值，对于 int 数组来说，元素的零值是 0。
	var a [5]int
	fmt.Println("emp", a)

	a[1] = 1
	fmt.Println("set", a)
	fmt.Println("get", a[1])

	fmt.Println("len:", len(a))

	//声明并初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

}
