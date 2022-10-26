package main

import (
	"fmt"
	"strings"
	"time"
)
import "math/rand"

func main() {
	var t = time.Now()
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
	for _, str := range input {
		for _, r := range str {
			hash[r]++
		}
	}
	for k, v := range hash {
		fmt.Printf("%c: %d\n", k, v)
	}
	fmt.Println("time.Since():", time.Since(t))
}
