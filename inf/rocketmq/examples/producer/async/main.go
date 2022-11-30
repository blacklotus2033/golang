package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const Topic = "testTopic"
const NamesrvAddr = "172.29.166.213:9876"
const Times int = 1e4

func main() {
	// #0 时间打印
	start := time.Now()
	// #1 创建并启动producer
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{NamesrvAddr}), //1)可以直接写 type NamesrvAddr []string
		//producer.WithRetry(10),
	)
	if err := p.Start(); err != nil {
		fmt.Println("p.Start() failed:", err)
		os.Exit(1)
	}
	// #2 启动循环 并发/异步地 发送消息
	var wg = new(sync.WaitGroup)
	var success uint64
	var cb = func(ctx context.Context, result *primitive.SendResult, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Println("p.SendAsync failed:", err)
			return
		}
		//success++
		atomic.AddUint64(&success, 1) //2)必须要做原子操作
		//fmt.Printf("p.SendAsync ok, success:%d, result:%s\n", success, result.String())
		//wg.Done() //1)不能在第一个if中return，又必须最后执行。改用defer。
	}
	for i := 0; i < Times; i++ {
		message := primitive.NewMessage(Topic, []byte("order:"+strconv.Itoa(i)))
		wg.Add(1)
		p.SendAsync(context.Background(), cb, message)
	}
	// #3 关闭producer
	wg.Wait()
	fmt.Println("After wg.Wati() success:", success)
	if err := p.Shutdown(); err != nil {
		fmt.Println("p.Shutdown() failed:", err)
		os.Exit(1)
	}
	fmt.Printf("Async Done. Times:%d Time elapsed:%s\n", Times, time.Since(start))
}
