package main

import "fmt"

// Constantes No Tipadas
//const Pi = 3.14159

// Constantes Tipadas
const Pi float32 = 3.14159

func getTypedConstants() {
	var number float64 = float64(Pi)

	fmt.Println(number)
}
