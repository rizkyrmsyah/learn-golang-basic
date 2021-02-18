package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println("time is", time.Now())

	fmt.Println("Your lucky number is", rand.Intn(100))
}
