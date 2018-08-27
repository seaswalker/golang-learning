package main

import (
	"strings"
	"fmt"
    "flag"
)

var n = flag.Bool("n", false, "omit trailing newline")
var seq = flag.String("s", " ", "separator")

// 测试指针的使用
func echo4() {
    // 解析自定义参数
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *seq));
    if !*n {
        fmt.Println()
    }
}
