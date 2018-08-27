package main

import (
	"fmt"
)

func testNew() {
    n := new(int);
    fmt.Println(*n)
    
    *n = 2;
    fmt.Println(*n)
}