package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg *sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg = new(sync.WaitGroup)

	for i := 1; i <= 1258; i++ {
		wg.Add(1)
		go printNum(i)
	}

	wg.Wait()
}

func printNum(a int) {
	fmt.Println(a)
	wg.Done()
}
