package main

import (
	"context"
	"fmt" //1)标准库Standard Library
	//
	"execlient/pb" //2)本地库Local
	//
	"google.golang.org/grpc" //3)第三方包
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Exe Grpc Client Starting...")
	//
	cc, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc.Dial err!")
		panic(err)
	}
	defer cc.Close()
	//
	{
		kc := pb.NewKnownClient(cc) //symmetry.1)
		//c := context.WithValue(context.Background(), "User-Agent", "My Exe Client")
		//c, _ := context.WithTimeout(context.Background(), time.Millisecond*500)
		c := context.Background()
		//fmt.Println("res.Root of 23.1:", RemoteSqrt(kc, c, 23.1))
		fmt.Println("greetings:", RemoteKnowFrom(kc, c))
	}
}

func RemoteSqrt(kc pb.KnownClient, c context.Context, input float64) float64 { //1)第一版简单封装
	req := &pb.SqrtReq{Input: input}
	res, err := kc.Sqrt(c, req)
	if err != nil {
		fmt.Println("kc.Sqrt err!")
		panic(err)
	}
	return res.Root
}

func RemoteKnowFrom(kc pb.KnownClient, c context.Context) string {
	req := &pb.KnowFromReq{}
	res, err := kc.KnowFrom(c, req)
	//res, err := kc.KnowFrom(c, nil) //panic: rpc error: code = Internal desc = grpc: error while marshaling: proto: Marshal called with nil
	if err != nil {
		fmt.Println("kc.KnowFrom err!")
		panic(err)
	}
	return res.Greetings
}
