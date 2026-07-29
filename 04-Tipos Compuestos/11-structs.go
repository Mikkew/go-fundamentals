package main

import "fmt"

func structsExample() {

	type person struct {
		name string
		age  int
		pet  string
	}

	var emiliano person

	// Struct literal
	fernando := person{}

	fmt.Println("Inicialmente:", emiliano, fernando)

	// Asignación de valores a los campos de la struct
	emiliano = person{
		name: "Emiliano",
		age:  20,
		pet:  "Perro",
	}

	fernando = person{
		name: "Fernando",
		age:  40,
		pet:  "Perros",
	}

	fmt.Println("Después de asignar valores:", emiliano, fernando)

	// Actualización de los campos de la struct
	emiliano.name = "Emiliano Andres"
	fmt.Println("Actualizacion 1", emiliano.name, fernando.name)
}
