package main

import (
	"fmt"
	"runtime"
	"sync"
)

const LoopTimes = int(1e9)
const GoMaxProcs = int(22222) //Trick)因为我的程序只有2个go程所以这里超过2是不会提升速度和占用CPU的。

func Ambiguous() int {
	var cnt int
	var wg = new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		for i := 0; i < LoopTimes; i++ {
			cnt *= 3
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < LoopTimes; i++ {
			cnt++
		}
		wg.Done()
	}()
	wg.Wait()
	return cnt
}

func main() {
	runtime.GOMAXPROCS(GoMaxProcs)
	for i := 1; i <= 999; i++ {
		fmt.Println("#", i, ":", Ambiguous())
	}
}
