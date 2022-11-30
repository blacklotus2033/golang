package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	go A(context.WithValue(context.Background(), "1", 1), 1)
	go A(context.WithValue(context.Background(), "2", 2), 2)
	time.Sleep(time.Second * 1)
}

func A(ctx context.Context, val int) {
	fmt.Println("A")
	B(ctx, val)
}
func B(ctx context.Context, val int) {
	fmt.Println("B")
	C(val)
}
func C(val int) {
	fmt.Println("C")
}
