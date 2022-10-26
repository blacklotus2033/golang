package main

import "fmt"

func main() {
	var in ChanInt = make(chan int)
	var out = in.GoPrint("cat").GoPrint("dog").GoPrint("fish")
	for i := 0; i < 10000; i++ {
		in <- i
		<-out
	}
}

type ChanInt chan int

func (in ChanInt) GoPrint(content string) ChanInt {
	var out = make(chan int)
	go func() {
		for {
			a := <-in
			fmt.Println(content, a)
			out <- a
		}
	}()
	return out
}
