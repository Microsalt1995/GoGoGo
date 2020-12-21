package main

import (
	"fmt"
	"math"
	"testing"
)

const s string = "constant"

func TestConstant(t *testing.T) {
	//const 用于声明一个常量。
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
