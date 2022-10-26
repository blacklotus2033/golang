package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type PairGen[T constraints.Float | constraints.Integer] struct {
	A, B T
}

func (v PairGen[T]) Get() (T, T) {
	return v.A, v.B
}

type Pointer PairGen[float32] //1)这种写法会使你的泛型失去意义
type Complex PairGen[int8]

//func SquareNested[T PairGen[U], U constraints.Float | constraints.Integer](from, to T) U { //2.a)错误的写法
//	from.Get()
//	return U(1)
//}

func SquareNested[T constraints.Float | constraints.Integer](from, to PairGen[T]) T { //2.b)正确的写法
	//fmt.Println(from.A) //3)只是goland提示不了，可以直接使用from的变量
	var da, db = from.A - to.A, from.B - to.B
	//var fa, fb = from.Get()
	//var ta, tb = to.Get()
	//var da, db = fa - ta, fb - tb
	return da*da + db*db
}

func PrintType(a any) { //1)不是泛型函数，仅仅是用了interface{}而已
	fmt.Printf("%v %T\n", a, a)
}

func main() {
	PrintType(SquareNested(PairGen[uint32]{1, 2}, PairGen[uint32]{4, 5}))
	var p = Pointer{1e2, 3e4}
	fmt.Println(p.A, p.B)
	//fmt.Println(p.Get()) //1)用不了。因为Pointer不会有PairGen[T]的泛型方法（generic method）。
	//SquareNested(Pointer{1e2, 3e4}, Pointer{4e5, 6e6}) //2) Pointer不再是PairGen[T]，这个其实在非泛型的情况下也是如此。
	//SquareNested(Complex{12, 34}, Complex{45, 66})
}
