package main

import (
	"google.golang.org/grpc"
	"grpc_server/myserver"
	"grpc_server/pb"
	"net"
)

func main() {
	s := grpc.NewServer() //1)这个和Client不是对称的
	{
		mks := new(myserver.MyKnownServer)
		pb.RegisterKnownServer(s, mks) //symmetry.1)
	}
	//
	l, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	//
	s.Serve(l)
}
