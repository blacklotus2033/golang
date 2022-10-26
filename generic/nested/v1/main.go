package main

import "golang.org/x/exp/constraints"

type PairGen[T constraints.Float | constraints.Integer] struct {
	A, B T
}

func (t PairGen[T]) GetA() T {
	return t.A
}

func (t PairGen[T]) GetB() T {
	return t.B
}

func SquareNested[T PairGen[U], U constraints.Float | constraints.Integer](from, to T) U {
	return from.GetA()
}

func main() {

}
