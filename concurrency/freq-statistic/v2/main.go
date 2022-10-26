package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)
import "math/rand"

func main() {
	var t = time.Now()
	runtime.GOMAXPROCS(1)
	fmt.Println("time.Since():", time.Since(t))
	var input = make([]string, 9999)
	var builder strings.Builder
	for i := 0; i < 9999; i++ {
		builder.Reset()
		for j := 0; j < 9999; j++ {
			builder.WriteRune(rune('a' + rand.Int()%26))
		}
		input[i] = builder.String()
	}
	fmt.Println("time.Since():", time.Since(t))
	//
	var hash = make(map[rune]int)
	var subs = make(chan map[rune]int)
	var wg = new(sync.WaitGroup)
	//var subs = make(chan map[rune]int, 100)
	for _, str := range input {
		//go func(str string) { //Trick.1)这是一个会直接报fatal error的一个写法。
		//	for _, r := range str {
		//		hash[r]++
		//	}
		//}(str)
		wg.Add(1)
		go func() { //1)直接闭包输入str和输出chan，待改为函数
			var sub = make(map[rune]int)
			for _, r := range str {
				sub[r]++
			}
			subs <- sub
			wg.Done()
		}()
	}
	go func() { //2)回收subs
		wg.Wait()
		close(subs)
	}()

	for sub := range subs {
		for k, v := range sub {
			hash[k] += v //3)本程序的核心就是：map-reduce
		}
	}
	var sum int
	for k, v := range hash {
		fmt.Printf("%c: %d\n", k, v)
		sum += v
	}
	fmt.Println("sum:", sum)
	fmt.Println("time.Since():", time.Since(t))
}
