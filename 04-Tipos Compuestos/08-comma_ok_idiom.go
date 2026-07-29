package main

import "fmt"

func commaOkIdiom() {
	myMap := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}

	// Comma ok idiom
	// Valor, valor booleano, si esta presente el valor en el mapa
	value, ok := myMap["Hola"]
	fmt.Println(value, ok)
}
