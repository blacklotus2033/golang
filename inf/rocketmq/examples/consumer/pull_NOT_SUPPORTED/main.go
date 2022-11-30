package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"os"
	"sync/atomic"
	"time"
)

const Topic = "testTopic"
const NamesrvAddr = "172.22.87.115:9876"
const ConsumerGroupName = "myCG"
const PullSleepTime = time.Second * 1
const PullNumbers = 100 //1)如果大于32是没有用的，没有找到限制的参数

// #0 时间打印
var start = time.Now()
var success uint64

/*
*
!!! 经试验整个PullConsumer是用不了的，无法PersistOffset()
*/
func main() {
	rlog.SetOutputPath("mq.log")
	// #1 创建并启动consumer
	//c, err := rocketmq.NewPullConsumer( //1)根据代码文档，暂未实现
	//The PullConsumer has not implemented completely, if you want have an experience of PullConsumer, you could use
	//// consumer.NewPullConsumer(...), but it may changed in the future.
	c, err := consumer.NewPullConsumer(
		consumer.WithNameServer([]string{NamesrvAddr}),
		consumer.WithGroupName(ConsumerGroupName),
	)
	if err != nil {
		fmt.Println("NewPullConsumer() failed:", err)
		os.Exit(1)
	}
	//if err := c.Subscribe(Topic, consumer.MessageSelector{}); err != nil {
	//	fmt.Println("c.Subscribe() failed:", err)
	//	os.Exit(1)
	//}
	if err := c.Start(); err != nil {
		fmt.Println("c.Start() failed:", err)
		os.Exit(1)
	}
	// #2 启动timer协程持久化offset 和 启动永久循环协程pull消息
	for i := 0; i < 100; i++ {
		if err := pull(c); err != nil {
			fmt.Println("pull_NOT_SUPPORTED() err:", err)
		}
	}
	if err := c.PersistOffset(context.TODO()); err != nil {
		fmt.Println("c.PersistOffset() err:", err)
	}
	// #3 关闭consumer
	if err := c.Shutdown(); err != nil {
		fmt.Println("c.Shutdown() failed:", err)
		os.Exit(1)
	}
	fmt.Printf("PullConsumer Done Time elapsed:%s\n", time.Since(start))
}

// #2.1 业务pull逻辑
func pull(c consumer.PullConsumer) error {
	res, err := c.Pull(context.TODO(), Topic, consumer.MessageSelector{}, PullNumbers)
	if err != nil {
		return fmt.Errorf("error doing something: %w", err) //1)Wrap
	}
	switch res.Status {
	case primitive.PullFound:
		break
	case primitive.PullNoNewMsg, primitive.PullNoMsgMatched:
		time.Sleep(PullSleepTime)
		return nil
	default:
		time.Sleep(PullSleepTime)
		return fmt.Errorf("wrong c.Pull() res.Status: %d", res.Status)
	}
	msgs := res.GetMessageExts()
	for _, msg := range msgs {
		atomic.AddUint64(&success, 1) //2)必须要做原子操作
		fmt.Printf("pull_NOT_SUPPORTED() at: %v success:%d queueId:%d QueueOffset:%d OffsetMsgId:%s msg.Body:%s\n", time.Since(start), success, msg.Queue.QueueId, msg.QueueOffset, msg.OffsetMsgId, msg.Body)
	}
	if err := c.UpdateOffset(msgs[0].Queue, res.NextBeginOffset); err != nil {
		fmt.Println("c.UpdateOffset() err:", err)
	}
	if err := c.PersistOffset(context.TODO()); err != nil {
		fmt.Println("c.PersistOffset() err:", err)
	}
	return nil
}
