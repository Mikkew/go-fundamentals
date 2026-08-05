package main

import "fmt"

func put(myMap map[string]int, key string, value int) {
	myMap[key] = value
}

func main() {
	var myMap map[string]int // nil
	fmt.Println("leer nil map:", myMap["x"])

	// myMap["a"] = 1 //panic
	myMap = make(map[string]int)
	fmt.Println("Mapa vacío: ", myMap == nil)

	put(myMap, "a", 1)
	put(myMap, "b", 2)
	fmt.Println("Mapa: ", myMap)
}
