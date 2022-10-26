package main

import (
	"fmt"
	"runtime"
	"time"
)

const LoopTimes = 999999

func main() { //1)这个程序是有二义性的，即便是单核CPU
	runtime.GOMAXPROCS(1)

	var cnt int
	go func() {
		for i := 0; i < LoopTimes; i++ {
			cnt *= 3
		}
	}()
	go func() {
		for i := 0; i < LoopTimes; i++ {
			cnt++
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("cnt:", cnt)
}
