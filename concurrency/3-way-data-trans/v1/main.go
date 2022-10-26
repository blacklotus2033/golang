package main

import (
	"fmt"
	"sync"
)

var g int
var wg *sync.WaitGroup

func MemWg() {
	g = 1
	wg.Done()
	//....
}

func Chan(ret chan<- int) {
	ret <- 1
	//....
}

func ChanPackaged() <-chan int {
	var ret = make(chan int)
	go func() {
		ret <- 1
	}()
	return ret
}

func main() {
	{ //1)
		wg = new(sync.WaitGroup)
		wg.Add(1)
		MemWg()
		wg.Wait()
		fmt.Println("g:", g)
	}
	{ //2)
		var ret = make(chan int)
		go Chan(ret)
		fmt.Println("<-ret:", <-ret)
	}
	{ //3)
		fmt.Println("<-ChanPackaged():", <-ChanPackaged())
	}
}
