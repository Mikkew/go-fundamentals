package main

import "fmt"

func sliceOfSlices() {
	x := []string{"a", "b", "c", "d", "e"}
	y := x[0:2]
	z := x[1:3]
	d := x[1:4]
	// Copia de memoria
	//e := x[:]
	// Creación de un nuevo slice con make y copia de memoria
	e := make([]string, 5)
	copyE := copy(e, x)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e, copyE)

	e[0] = "x"
	fmt.Println(x)
	fmt.Println(e)
}
