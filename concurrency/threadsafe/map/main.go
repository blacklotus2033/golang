package main

import (
	"fmt"
	"time"
)

func main() {
	//var hash map[int]int
	var hash = make(map[int]int)
	//var hash = new(sync.Map)
	for i := 0; i <= 9; i++ {
		//go func() {
		//	hash[i] = i / i //Trick.1)当闭包i的时候，运行时，i已经新值了。不会触发“integer divide by zero”。
		//}()
		//go func(i int) {
		//	defer SimpleRecover()
		//	hash[i] = i / i //1)故意引发“integer divide by zero”的panic。
		//}(i)
		go func(i int) {
			defer SimpleRecover()
			hash[i] = i //2)故意触发fatal error非panic
		}(i)
		//go func() {
		//	hash.Store(i, i)
		//}
	}
	time.Sleep(2 * time.Second)
	fmt.Println("len of hash:", len(hash))
	fmt.Println("hash:", hash)
}

func SimpleRecover() {
	if r := recover(); r != nil {
		fmt.Println("recover():", r)
	}
}
