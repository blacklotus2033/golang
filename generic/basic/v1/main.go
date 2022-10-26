package main

import (
	"fmt"
)

func Sum[T int | float64 | string](arr ...T) T {
	var sum T
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

type MyInt int

func SumUnderlying[T ~int | ~float64 | ~string](arr ...T) T {
	var sum T
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1e3, 2e4, 3e5))
	fmt.Println(Sum("a", "bb", "ccc"))
	//fmt.Println(Sum()) //1)编译错：cannot infer T
	//fmt.Println(Sum(MyInt(1), MyInt(1), MyInt(1)))
	fmt.Println(SumUnderlying(MyInt(1), MyInt(1), MyInt(1)))
}
