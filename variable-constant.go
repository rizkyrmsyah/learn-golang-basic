package main

import (
	"fmt"
)

func main() {
	const Pi = 3.14
	var my_name = "rizky"
	fmt.Println(Pi, my_name)

	country, city, province := "indonesia", "bandung", "JAWA BARAT"
	fmt.Println(country, city, province)

	var age, weight float64
	age, weight = 24, 66
	fmt.Println(age, weight)
}
