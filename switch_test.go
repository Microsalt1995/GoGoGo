package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}

}
