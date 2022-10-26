package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	WorkerNum = 4
	ProcNum   = WorkerNum + 5
	TargetSum = 1e8
	ChanBuf   = 1e4
)

var sum int64
var mu sync.Mutex

func ByMutex(wg *sync.WaitGroup) { //1)闭包sum，使用WorkerNum、TargetSum
	for i := 0; i < TargetSum/WorkerNum; i++ {
		mu.Lock()
		sum++
		mu.Unlock()
	}
	wg.Done()
}

func ByAtomic(wg *sync.WaitGroup) { //1)闭包sum，使用WorkerNum、TargetSum
	for i := 0; i < TargetSum/WorkerNum; i++ {
		atomic.AddInt64(&sum, 1)
	}
	wg.Done()
}

func ByChanProduce(wg *sync.WaitGroup, ch chan<- int64) {
	for i := 0; i < TargetSum/WorkerNum; i++ {
		ch <- 1
	}
	wg.Done()
}

func ByChanConsume(ch <-chan int64) { //1)闭包sum
	for v := range ch {
		sum += v
	}
	fmt.Println("quit ByChanConsume!")
}

func main() {
	runtime.GOMAXPROCS(ProcNum)
	t := time.Now()
	//
	//
	wg := new(sync.WaitGroup)
	{
		fmt.Println("ByMutex")
		for i := 0; i < WorkerNum; i++ {
			wg.Add(1)
			go ByMutex(wg)
			//go ByAtomic(wg)
		}
		wg.Wait()
		fmt.Println("TargetSum:", TargetSum, "WorkerNum:", WorkerNum, "sum:", sum)
		fmt.Println("time.Since():", time.Since(t))
		t = time.Now()
	}
	{
		fmt.Println("ByAtomic")
		for i := 0; i < WorkerNum; i++ {
			wg.Add(1)
			//go ByMutex(wg)
			go ByAtomic(wg)
		}
		wg.Wait()
		fmt.Println("TargetSum:", TargetSum, "WorkerNum:", WorkerNum, "sum:", sum)
		fmt.Println("time.Since():", time.Since(t))
		t = time.Now()
	}
	{
		fmt.Println("ByChan")
		ch := make(chan int64, ChanBuf)
		go ByChanConsume(ch)
		for i := 0; i < WorkerNum; i++ {
			wg.Add(1)
			//go ByMutex(wg)
			//go ByAtomic(wg)
			go ByChanProduce(wg, ch)
		}
		wg.Wait()
		close(ch)
		fmt.Println("TargetSum:", TargetSum, "WorkerNum:", WorkerNum, "sum:", sum)
		fmt.Println("time.Since():", time.Since(t))
		t = time.Now()
	}
}
