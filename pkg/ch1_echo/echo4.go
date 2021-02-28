package main

import (
	"fmt"
	"os"
)

// Exercise 1.2
func main() {
	for index, arg := range os.Args[1:] {
		print := fmt.Sprintf("%d - %s", index, arg)
		fmt.Println(print)
	}
}
