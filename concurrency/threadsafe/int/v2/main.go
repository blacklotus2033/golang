package main

import (
	"fmt"
	"time"
)

func main() {
	var cnt int
	for i := 0; i < 9999; i++ {
		go func() { //1)go程模拟并发
			time.Sleep(time.Second)
			cnt++
		}()
	}
	time.Sleep(2 * time.Second) //2)等待是不可或缺的
	fmt.Println("cnt:", cnt)    //3)这个程序最严重的问题是：正确性得不到保证
}
