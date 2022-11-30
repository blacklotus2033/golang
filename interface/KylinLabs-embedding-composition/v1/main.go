package main

import "fmt"

func main() {
	var t = new(Teacher)
	t.Show()
}

type (
	People struct { //1)空结构
	}
	Teacher struct { //2)“嵌入embedding”
		People
	}
)

func (p *People) Show() {
	fmt.Println("Show")
	p.Show1()
}

func (p *People) Show1() {
	fmt.Println("Show1")
}

func (p *Teacher) Show1() {
	fmt.Println("Teacher Show1")
}
