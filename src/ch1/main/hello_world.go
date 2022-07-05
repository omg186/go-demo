package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world!")
	s := os.Args
	fmt.Printf("args: %v\n", s)
	if len(s) > 1 {
		fmt.Printf("s: %v\n", s[1])
	}
}
