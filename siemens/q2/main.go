package main

import "fmt"

type statement struct {
	cmd string
	params []string
	Children []statement
}

func main() {
	var input = "[when_system_start]\n" +
		"-[assign](age,15)\n" +
		"-[if](age, >=, 18)\n" +
		"--[print](You are an adult now)\n" +
		"-[else]\n" +
		"--[print](You are still young)"

	st: = statement{cmd: "", params: nil, Children: []statement{
		statement{cmd: "assign", params: []string{"age", "15"}},
		statement{cmd: "if", params: []string{"age", ">=", "18"}, Children: []statement{
			statement{cmd: "print", params: []string{""}},
			}
		},
		statement{cmd: "else", params: nil, Children: []statement{

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
