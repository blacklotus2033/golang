package main

import (
	"example.com/golang/rpc/protobuf/v1/pb"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	var co = &pba.CoffeeOrder{
		Uid:      10001,
		UserName: "潘达叔",
		Note:     "加冰，自带杯",
		Items: []*pba.CoffeeOrder_CoffeeItem{
			{
				Cid:     20003,
				Name:    "西班牙拿铁",
				Quality: 1,
				Each:    22.2,
				Price:   22.2,
			},
			{
				Cid:     20004,
				Name:    "椰菠冷萃",
				Quality: 1,
				Each:    33.3,
				Price:   33.3,
			},
		},
		TotalPrice: 55.5,
	}
	marshalled, err := proto.Marshal(co)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshalled:%v\n", marshalled)
	var unmarshalled = new(pba.CoffeeOrder)
	err = proto.Unmarshal(marshalled, unmarshalled)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshalled:%#v\n", unmarshalled)
}
