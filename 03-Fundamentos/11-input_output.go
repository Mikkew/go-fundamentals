package main

import "fmt"

func inputOutput() {
	var name string
	var age int

	// Antes de leer los valores, imprimimos los valores iniciales de las variables y sus direcciones de memoria
	fmt.Println(name, age)
	fmt.Println(&name)
	fmt.Println(&age)

	fmt.Println("Ingresa tu nombre: ")
	//& es para pasar la referencia de la variable, ya que Scan necesita modificar el valor de la variable
	// apuntando a la dirección de memoria de la variable name
	fmt.Scan(&name)

	fmt.Println("Ingresa tu edad: ")
	fmt.Scan(&age)

	fmt.Printf("Hola %s, tienes %d años.\n", name, age)

	// Después de leer los valores, imprimimos nuevamente los valores de las variables y sus direcciones de memoria
	fmt.Println(name, age)
	fmt.Println(&name)
	fmt.Println(&age)
}
