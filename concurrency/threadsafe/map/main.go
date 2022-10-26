package main

import (
	"fmt"
	"time"
)

func main() {
	var arr []string
	for i := 0; i < 9999; i++ {
		go func() { //1)go程模拟并发
			//time.Sleep(time.Second)
			arr = append(arr, "somewhat")
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("len of arr:", len(arr))
}
