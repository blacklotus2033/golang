package main

import "fmt"

type Tools interface {
	Show()
	Show1()
}

func Process(t Tools) {
	t.Show()
	t.Show1()
}

type MyTools struct{}

func (p *MyTools) Show() {
	fmt.Println("MyTools Show")
}
func (p *MyTools) Show1() {
	fmt.Println("MyTools Show1")
}

type YourTools struct{}

func (p *YourTools) Show() {
	fmt.Println("YourTools Show")
}
func (p *YourTools) Show1() {
	fmt.Println("YourTools Show1")
}

func main() {
	var m, y = new(MyTools), new(YourTools)
	Process(m)
	Process(y)
}
