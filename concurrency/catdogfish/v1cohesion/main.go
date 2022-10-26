package main

import "fmt"

func main() {
	var in = make(chan int)
	catOut := GoPrint("cat", in)
	dogOut := GoPrint("dog", catOut)
	fishOut := GoPrint("fish", dogOut)
	for i := 0; i < 100; i++ {
		in <- i
		<-fishOut
	}
}

func GoPrint(content string, in <-chan int) <-chan int {
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
