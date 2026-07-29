package main

import "fmt"

func sliceFunctions() {
	// var mySlice []int
	// fmt.Println(mySlice == nil)

	// mySlice := []int{1, 2, 3, 4, 5}
	// fmt.Println(mySlice)

	// len: Para obtener la longitud de un slice, se utiliza la función len(). Esta función devuelve el número de elementos presentes en el slice.
	// fmt.Println(len(mySlice))

	// append: Para agregar elementos a un slice, se utiliza la función append(). Esta función toma un slice existente y uno o más elementos adicionales, y devuelve un nuevo slice que contiene todos los elementos.
	// mySlice = append(mySlice, 10, 12, 30, 12, 40, 50, 12, 3, 12, 35, 12, 13)
	// fmt.Println(mySlice)

	// Capacity: Para obtener la capacidad de un slice, se utiliza la función cap(). La capacidad representa el número máximo de elementos que el slice puede contener antes de necesitar una reasignación de memoria.
	// fmt.Println(mySlice, len(mySlice), cap(mySlice))
	// fmt.Println(len(mySlice) == cap(mySlice))

	makeSlice := make([]int, 0, 10)
	fmt.Println(makeSlice)

	makeSlice = append(makeSlice, 10, 20, 30, 40, 10, 20, 30, 40, 10, 20, 30, 40)
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice), cap(makeSlice))

	// Vaciar slice
	clear(makeSlice)
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice), cap(makeSlice))
}
