package main

import (
	"fmt"
)

func main() {
	var in = make(chan int)
	catOut := GoPrint("cat", in)
	dogOut := GoPrint("dog", catOut)
	fishOut := GoPrint("fish", dogOut)
	for i := 0; i < 10000; i++ {
		in <- i
		<-fishOut
	}
	close(in)
	//time.Sleep(time.Second)
	//fmt.Println("reading fishOut:", <-fishOut)
	for range fishOut {
	}
}

func GoPrint(content string, in <-chan int) <-chan int {
	var out = make(chan int)
	go func() {
		for a := range in {
			fmt.Println(content, a)
			out <- a
		}
		fmt.Println("Content:", content, " input channel closing!")
		close(out)
	}()
	return out
}
