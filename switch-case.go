package main

import "fmt"

func main() {
	value := 100

	switch value {
	case 100:
		fmt.Println("true")
	default:
		fmt.Println("false")
	}
}
