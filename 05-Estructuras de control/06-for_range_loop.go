package main

import "fmt"

func forRangeLoop() {
	evenNumbers := []int{2, 4, 6, 8, 10, 12}

	//Itera index y valor
	for index, value := range evenNumbers {
		//Iterar el index y el valor
		fmt.Println(index, value)

		fmt.Println(value)
	}

	// Posiciones
	for k := range evenNumbers {
		fmt.Println(k)
	}

	// Valores
	for _, value := range evenNumbers {
		// Iterar el index y el valor
		//fmt.Println(index, value)

		fmt.Println(value)
	}
}
