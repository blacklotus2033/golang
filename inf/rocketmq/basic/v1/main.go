package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	endPoint := []string{"172.19.251.168:9876"}
	p, err := rocketmq.NewProducer(
		//producer.WithNameServer(endPoint),
		producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2),
		//producer.WithGroupName("GID_xxxxxx"),
	)
	if err != nil {
		fmt.Println("NewProducer err:", err)
	}
	err = p.Start()
	if err != nil {
		fmt.Println("Start err:", err)
	}
	var msgs []*primitive.Message
	msgs = append(msgs, primitive.NewMessage(
		"test",
		[]byte("SendSync Hello RocketMQ Go Client!"),
	))
	res, err := p.SendSync(context.Background(), msgs...)
	a, b := primitive.NewMessage("test", []byte("a")), primitive.NewMessage("test", []byte("b"))
	p.SendSync(context.Background(), a)
	p.SendSync(context.Background(), a, b)
	//var arr [2]*primitive.Message
	//arr[0], arr[1] = a, b
	//p.SendSync(context.Background(), arr...)

	if err != nil {
		fmt.Println("SendSync err:", err)
	}
	fmt.Println("res:", res)

	//err = p.SendOneWay(context.Background(), primitive.NewMessage(
	//	"test",
	//	[]byte("SendOneWay Hello RocketMQ Go Client!"),
	//))
	//if err != nil {
	//	fmt.Println("SendOneWay err:", err)
	//}

	p.Shutdown()
}
