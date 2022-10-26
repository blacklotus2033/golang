package main

import (
	"bytes"
	"fmt"
	"golang.org/x/time/rate"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Simple recover():", r)
		}
	}()

	var c = make(chan int)
	<-c //Trick)fatal error测试：all goroutines are asleep - deadlock!

	var a int
	println(a)
	//fmt.Println(a / a) //Trick)panic测试

	b := bytes.Buffer{}
	//println(b) //Trick)会直接报CE（编译错）：illegal types for operand: print -- Compilation finished with exit code 2
	fmt.Println(b)

	limiter := rate.NewLimiter(10, 100)
	fmt.Println(limiter.Allow())
}
