package main

import "fmt"

type Statement struct {
	Cmd      string
	Params   []string
	Children []Statement
}

func main() {
	var input = "[when_system_start]\n" +
		"-[assign](age,15)\n" +
		"-[if](age, >=, 18)\n" +
		"--[print](You are an adult now)\n" +
		"-[else]\n" +
		"--[print](You are still young)"

	st: = Statement{Cmd: "", Params: nil, Children: []Statement{
		Statement{Cmd: "assign", Params: []string{"age", "15"}},
		Statement{Cmd: "if", Params: []string{"age", ">=", "18"}, Children: []Statement{
			Statement{Cmd: "print", Params: []string{""}},
			}
		},
		Statement{Cmd: "else", Params: nil, Children: []Statement{

		}},
	}}

	age := 15
	if age >= 18 {
		fmt.Println("")
	} else {

	}
}

cmd2format := map[string]string {
	"assign": "%s := %s",
	"if": "if %s {%s}", //
	"else": "else {%s}"
}
