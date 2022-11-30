package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"time"
)

const (
	Topic             = "TopicTest"
	NamesrvAddr       = "172.21.87.234:9876"
	ProducerGroupName = "myPG"
	Loops             = 1000
	MessageBatchSize  = 1
	SendSyncSleepTime = time.Millisecond * 0
	MsgBodyFormat     = "From producer/batch/main.go i-in-loop:%d total:%d"
	MsgShardingKey    = "mySK"
)

func main() {
	endPoint := []string{NamesrvAddr}
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2),
		producer.WithGroupName(ProducerGroupName),
		producer.WithQueueSelector(producer.NewHashQueueSelector()), //1)加上这个后，ShardingKey才有效。
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	var total int
	for x := 0; x < Loops; x++ {
		var msgs []*primitive.Message
		for i := 0; i < MessageBatchSize; i++ {
			total++
			msg := primitive.NewMessage(Topic, []byte(fmt.Sprintf(MsgBodyFormat, i, total)))
			//msg = msg.WithShardingKey(MsgShardingKey)
			msg.WithShardingKey(MsgShardingKey)
			msgs = append(msgs, msg)
		}

		res, err := p.SendSync(context.Background(), msgs...)
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
		time.Sleep(SendSyncSleepTime)
	}

	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
