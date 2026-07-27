package main

import "fmt"

var age int

func declarations() {
	//Declaración de variables
	/*
			Declaracion explicita, aplicar cuando se quiere especificar el tipo de dato
			cuando quieras un valor inicial, puedes asignarlo en la misma línea de la declaración
		 	cuando estas a nivel de paquete, es decir, a fuera de las funciones, no puedes usar la declaración corta,
			ya que esta solo se puede usar dentro de las funciones
	*/
	//var age int = 30
	age = 30

	//Declaracion corta, aplicar cuando se quiere que el compilador infiera el tipo de dato
	name := "Miguel"

	fmt.Println(age)
	fmt.Println(name)
}
