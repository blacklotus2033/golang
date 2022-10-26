package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type Pair[T constraints.Float] struct {
	A, B T
}

type Point Pair[float32]
type Complex Pair[float64]

func (v Pair[T]) Distance() T {
	return T(math.Sqrt(float64(v.A*v.A + v.B*v.B)))
}

func main() {
	var p = Point{1, 2}
	var c = Complex{5, 3}
	fmt.Println(Pair[float32](p).Distance())
	fmt.Println(Pair[float64](c).Distance())
}
