package main

import "fmt"

func arrays() {
	// var number_list [3]int

	// Array literal
	// var numberList = [3]int{10, 20, 30}

	// Array literal simplificado
	var numberList = [...]int{10, 20, 30, 40}
	fmt.Println(numberList)

	numberList[0] = 50
	fmt.Println(numberList[2])
}
