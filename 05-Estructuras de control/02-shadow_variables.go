package main

import "fmt"

func shadowVariables() {
	// Scope de la variable aplica dentro de su bloque
	x := 10

	if x > 5 {
		fmt.Println(x)

		//El scope aplica dentro de su propio bloque y aplica como variable nueva
		//x := 5

		// Se reasigna el valor global de la variable
		x = 5

		fmt.Println(x)
	}

	fmt.Println(x)
}
