package main

import (
	"fmt"
	"os"
	"strings"
)

func echoMain() {
	stupidLoop()
	betterLoop()
	noLoop()
}

func stupidLoop() {
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
}

func betterLoop() {
	s, seq := "", ""

	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}

	fmt.Println(s)
}

func noLoop() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
