package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	AccLimit = int64(1e8)
	GoNum    = 4
	GoProcs  = 1
)

func main() {
	runtime.GOMAXPROCS(GoProcs)
	//
	var accumulator int64
	var mu = new(sync.Mutex)
	var startUs = time.Now().UnixMicro()
	//var startNs = time.Now().UnixNano()
	var change [GoNum][][2]int64
	var done = make(chan int)
	var add = func(mark int) {
		var last int64 = -1
		for {
			mu.Lock()
			if accumulator >= AccLimit {
				mu.Unlock()
				done <- mark
				break
			}
			//time.Sleep(time.Nanosecond * 1)
			if accumulator != last {
				//fmt.Printf("mark:%d time(us):%d diff:%d last:%d acc:%d\n", mark, time.Now().UnixMicro()-startUs, accumulator-last, last, accumulator)
				//fmt.Printf("mark:%d time(us):%d diff:%d last:%d acc:%d\n", mark, time.Now().UnixNano()-startNs, accumulator-last, last, accumulator)
				change[mark] = append(change[mark], [2]int64{time.Now().UnixMicro() - startUs, last})
				//change[mark] = append(change[mark], [2]int64{time.Now().UnixNano() - startNs, last})
			}
			accumulator++
			last = accumulator
			mu.Unlock()
		}
	}
	for i := 0; i < GoNum; i++ {
		go add(i)
	}
	select { //1)废话select
	case mark := <-done:
		fmt.Printf("#%d has done at %d us.\n", mark, time.Now().UnixMicro()-startUs)
	}
	fmt.Println("accumulator:", accumulator)

	for i := range change {
		//fmt.Printf("#%d : %v\n", i, change[i])
		fmt.Printf("#%d len: %d\n", i, len(change[i]))
	}
}
