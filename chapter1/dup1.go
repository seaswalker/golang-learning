package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup1Main() {

	// key: 每一行，value: 次数
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		text := input.Text()

		if text == "done" {
			break
		}

		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}
