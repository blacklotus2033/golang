package main

import (
	"fmt"
	"time"
)

const Seconds = 1

func main() {
	start := time.Now()
	t := time.NewTimer(time.Second * Seconds)
	var cnt int
	for ; true; <-t.C { //1)相当于是 do while
		t.Reset(time.Second * Seconds)
		cnt++
		fmt.Printf("cnt = %d at elapsed:%v\n", cnt, time.Since(start))
	}
}
