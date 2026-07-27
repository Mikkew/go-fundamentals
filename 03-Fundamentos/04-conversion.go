package main

import "fmt"

func getConversion() {
	var number1 int = 10
	var number2 float64 = 3.5

	total := float64(number1) + number2 // Convert int to float64 for addition
	fmt.Printf("Total: %v \n", total)
}
