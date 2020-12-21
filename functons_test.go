package main

import (
	"fmt"
	"testing"
)

//这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	return a + b
}

//多返回值
func vals() (int, int) {
	return 3, 7
}

//当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
	return a + b + c
}

//变参函数
//这个函数接受任意数量的 int 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//Go 支持匿名函数， 并能用其构造 闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func TestFunctions(t *testing.T) {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	fmt.Println()

	sum(1, 2)
	sum(1, 2, 3)

	//如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	fmt.Println()

	//我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println()
	fmt.Println(fact(7))
}
