package main

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {
	//var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//声明后却没有给出对应的初始值时，变量将会初始化为 零值 。
	var e int
	fmt.Println(e)

	//:= 语法是声明并初始化变量的简写， 例如 var f string = "short" 可以简写为右边这样。
	f := "short"
	fmt.Println(f)

}
