package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

const (
	Topic             = "TopicTest"
	NamesrvAddr       = "172.25.98.139:9876"
	ConsumerGroupName = "myCG"
)

var (
	EmptySelector = consumer.MessageSelector{}
)

func main() {
	// #0 RocketMQ日志重定向
	//rlog.SetOutputPath("./mq.log")

	// #S.1 开启信号拦截（signal interception）
	sigs := make(chan os.Signal)
	//signal.Notify(sigs, os.Interrupt) //两种写法都可以
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //

	var c rocketmq.PushConsumer
	// #1 创建PushConsumer
	{
		var err error
		if c, err = rocketmq.NewPushConsumer(
			consumer.WithGroupName(ConsumerGroupName),
			consumer.WithNameServer([]string{NamesrvAddr}),
			//consumer.WithConsumerOrder(true),
			//consumer.WithConsumeConcurrentlyMaxSpan(1),
			//consumer.WithConsumeMessageBatchMaxSize(1),
			//consumer.WithPullBatchSize(1),
		); err != nil {
			fmt.Println("rocketmq.NewPushConsumer() error:", err)
			os.Exit(1)
		}
	}

	// #2 进行订阅，并指定回调函数
	{
		var mu sync.Mutex
		var callId int
		var CallId = func() int {
			mu.Lock()
			defer mu.Unlock()
			callId++
			return callId
		}
		var concurrency, max int64 //max的128测出来可能是：4x32
		// 回调函数原型要求：f func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)
		var consume = func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			//统计.1) 完成并发统计
			atomic.AddInt64(&concurrency, 1)
			var con = atomic.LoadInt64(&concurrency)
			if max < con { //一个不安全的统计
				max = con
			}
			fmt.Println("max concurrency:", max, "consume concurrency:", con)
			defer atomic.AddInt64(&concurrency, -1)

			//统计.2) 获取当前callId从1开始
			id := CallId()
			//fmt.Println("CallId() id:", id)

			//业务.1) 循环处理收集参数msgs
			for i, msg := range msgs {
				const processFormat = "concurrency:%d id:%d i:%d msg.GetShardingKey():%s queueId:%d msg.Body:%s\n"
				fmt.Printf(processFormat, concurrency, id, i, msg.GetShardingKey(), msg.Queue.QueueId, msg.Body)
			}
			return consumer.ConsumeSuccess, nil
		}
		if err := c.Subscribe(Topic, EmptySelector, consume); err != nil {
			fmt.Println("c.Subscribe() error:", err)
			os.Exit(1)
		}
	}
	// #3 启动PushConsumer
	if err := c.Start(); err != nil {
		fmt.Println("c.Start() error:", err)
		os.Exit(1)
	}

	// #S.2 等待信号
	sig := <-sigs
	fmt.Printf("signal:%s received.\n", sig)
	// #4 关闭PushConsumer
	if err := c.Shutdown(); err != nil {
		fmt.Println("c.Shutdown() error:", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
