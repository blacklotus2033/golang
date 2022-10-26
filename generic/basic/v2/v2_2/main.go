package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type Pair[T constraints.Float | constraints.Integer] struct {
	A, B T
}

func (v Pair[T]) Distance() T {
	return T(math.Sqrt(float64(v.A*v.A + v.B*v.B)))
	//var fa, fb = float64(v.A), float64(v.B) //2.a)改进型。
	//return T(math.Sqrt(fa*fa + fb*fb))
}

func main() {
	var p = Pair[float32]{12, 5.5}
	var c = Pair[float64]{3, 4}
	var ii = Pair[uint64]{1.2e9, 5e8}
	var io = Pair[uint32]{1.2e9, 5e8} //2)在做v.A*v.A的时候越界了。改进后可以解决。
	var ir = Pair[int]{100, 100}      //1)有可能强转T()的泛型逻辑并不一定符合我们的需求。
	fmt.Println(p.Distance())
	fmt.Println(c.Distance())
	fmt.Println(ii.Distance())
	fmt.Println(io.Distance())
	fmt.Println(ir.Distance())
}
