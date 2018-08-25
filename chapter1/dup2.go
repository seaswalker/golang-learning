package main

import (
	"bufio"
	"fmt"
	"os"
)

// 从给定的输入(文件)统计重复行
func dup2Main(file string) {
	counts := make(map[string]int)

	f, err := os.Open(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v,\n", err)
	}

	countLines(f, counts)
	f.Close()

	// 打印统计结果
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
