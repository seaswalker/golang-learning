package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 一次读入所有的文件内容
func dup3(file string) {
	counts := make(map[string]int)

	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v,\n", err)
	}

	// Windows上使用\r\n才是正确的结果
	for _, line := range strings.Split(string(data), "\r\n") {
		counts[line]++
	}

	// 打印统计结果
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}
