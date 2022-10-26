package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	co := CoffeeOrder{
		Uid:      1001,
		UserName: "陈潘达",
		Note:     "少糖，少冰",
		Items: []CoffeeItem{
			{20001,
				"馥芮白",
				3,
				5.11,
				15.33,
			},
			{20002,
				"香橙Dirty",
				1,
				7.00,
				7.00,
			},
		},
		TotalPrice: 22.33,
	}
	var marshalled, err = jsoniter.MarshalIndent(co, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshaled(string):%s\n\n", string(marshalled))
	fmt.Printf("Valid:%#v\n\n", jsoniter.Valid(marshalled))
	var unmarshalled CoffeeOrder
	err = jsoniter.Unmarshal(marshalled, &unmarshalled)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshaled:%#v\n\n", unmarshalled)
}
