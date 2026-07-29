package main

import "fmt"

func mapsDelete() {
	myMap := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}
	fmt.Println(myMap)

	delete(myMap, "Hola") // Eliminar un elemento del map
	clear(myMap)          // Limpiar todos los elementos

	fmt.Println(myMap)
	fmt.Println(len(myMap))
}
