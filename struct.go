package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  bool
}

func main() {
	fmt.Println(Person{"rizky", 25, true})
}
