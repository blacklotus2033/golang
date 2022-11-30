package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
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
	)
	if err := p.Start(); err != nil {
		fmt.Println("p.Start() failed:", err)
		os.Exit(1)
	}
	// #2 启动循环 同步地 发送消息
	for i := 0; i < 1e4; i++ {
		message := primitive.NewMessage(Topic, []byte("order:"+strconv.Itoa(i)))
		p.SendSync(context.Background(), message)
	}
	// #3 关闭producer
	if err := p.Shutdown(); err != nil {
		fmt.Println("p.Shutdown() failed:", err)
		os.Exit(1)
	}
	fmt.Printf("--COMPARE-- Sync Done. Times:%d Time elapsed:%s\n", Times, time.Since(start))
}
