package main

import (
	"fmt"
	"log"
	"regexp"
)

func run(input string) {
	re := regexp.MustCompile(`^\[(\w+)]\n.\[(\w+)]\((.*)\)$`)
	params := re.FindStringSubmatch(input)
	//for _, param := range params {
	//	fmt.Println(param)
	//}
	cmd, param := params[2], params[3]
	//fmt.Println(cmd)
	switch cmd {
	case "print":
		fmt.Println(param)
	default:
		log.Fatalln("Unknown cmd.")
	}
}

func main() {
	var input = "[when_system_start]\n-[print](hello,world)"
	run(input)
}
