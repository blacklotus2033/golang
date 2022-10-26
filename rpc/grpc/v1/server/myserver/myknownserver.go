package myserver

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

import "grpc_server/pb" //1)自己的包

type MyKnownServer struct {
	pb.UnimplementedKnownServer
}

func (p *MyKnownServer) Sqrt(_ context.Context, req *pb.SqrtReq) (*pb.SqrtRes, error) {
	if req.Input < 0 {
		return nil, errors.New("input invalid: Negative Number") //1)这里会返回nil，看client如何处理
	}
	res := new(pb.SqrtRes)
	res.Root = math.Sqrt(req.Input)
	return res, nil
}

func (p *MyKnownServer) KnowFrom(c context.Context, _ *pb.KnowFromReq) (*pb.KnowFromRes, error) {
	//1)gin框架中通过：ua2 := c.Request.Header.Get("User-Agent") 获取
	fmt.Printf("c:%#v\n  c.Value(\"User-Agent\"):%#v\n", c, c.Value("User-Agent"))
	ua, _ := c.Value("User-Agent").(string)
	//ua, ok := c.Value("User-Agent").(string)
	//if !ok {
	//	return nil, errors.New("context invalid: c.Value(\"User-Agent\").(string) failed")
	//}
	res := new(pb.KnowFromRes)
	res.Greetings = fmt.Sprintf("I known you are from UA:%s", ua)
	{ //2)超时实验
		time.Sleep(time.Millisecond * 1500)
	}
	select {
	case <-c.Done():
		fmt.Println("context already Done!")
	default:
		fmt.Println("After sleep 1.5 second.")
	}
	return res, nil
}
