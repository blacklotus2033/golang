package main

type CoffeeOrder struct {
	Uid        int
	UserName   string
	Note       string
	Items      []CoffeeItem
	TotalPrice float64
}

type CoffeeItem struct {
	Cid     int
	Name    string
	Quality int
	Each    float64
	Price   float64
}
