package main

import (
	"fmt"
	"testing"
)

func TestIf_else(t *testing.T) {
	//注意，在 Go 中，条件语句的圆括号()不是必需的，但是花括号{}是必需的。
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

}
