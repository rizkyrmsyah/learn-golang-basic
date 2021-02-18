package main

import (
	"fmt"
)

func main() {
	var result = add(100, 16)

	fmt.Println("result is", result)
}

func add(x int, y int) int {
	return x + y
}
