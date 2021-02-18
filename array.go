package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Muhamad"
	name[1] = "Rizky"
	name[2] = "Rahmansyah"

	fmt.Println(name[0], name[1], name[2])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
