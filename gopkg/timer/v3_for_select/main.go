package main

import (
	"fmt"
	"time"
)

const Seconds = 2

func main() {
	start := time.Now()
	t := time.NewTimer(time.Second * Seconds)
	var cnt int
	for {
		select {
		case <-t.C:
			t.Reset(time.Second * Seconds)
			cnt++
			fmt.Printf("cnt = %d at elapsed:%v\n", cnt, time.Since(start))
		}
	}
}
