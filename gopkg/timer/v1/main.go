package main

import (
	"fmt"
	"time"
)

const Seconds = 5

func main() {
	t := time.NewTimer(time.Second * Seconds)
	<-t.C
	fmt.Printf("After %d second(s).\n", Seconds)
}
