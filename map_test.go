package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	a, prs := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("a:", a)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
