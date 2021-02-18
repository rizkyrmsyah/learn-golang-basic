package main

import (
	"fmt"
)

func main() {
	a, b, c := multi("Hello", "bad", "niggas")

	fmt.Println(a, b, c)
}

func multi(x, y, z string) (string, string, string) {
	return x, y, z
}
