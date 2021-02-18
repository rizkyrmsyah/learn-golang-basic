package main

import "fmt"

func main() {

	// For is Go's "while"
	// At that point you can drop the semicolons: C's while is spelled for in Go.

	sum := 1
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}
